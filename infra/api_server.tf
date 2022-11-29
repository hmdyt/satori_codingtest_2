locals {
  envs = {
    "MYSQL_USER"     = var.mysql_user,
    "MYSQL_PASSWORD" = var.mysql_password,
    "MYSQL_DATABASE" = var.mysql_database,
    "MYSQL_NET"      = "unix",
    "MYSQL_HOST"     = "/cloudsql/${var.project_id}:${var.region}:${google_sql_database_instance.sql.name}"
    "NEXT_HOST"      = google_cloud_run_service.web.status[0].url
  }
}
data "archive_file" "api_archive_file" {
  type        = "zip"
  source_dir  = var.api_dir
  output_path = var.api_output_path
}

resource "google_storage_bucket" "bucket" {
  name     = "${var.project_id}-api"
  location = var.region
}

resource "google_storage_bucket_object" "zip" {
  name   = "${data.archive_file.api_archive_file.output_md5}.zip"
  bucket = google_storage_bucket.bucket.name
  source = data.archive_file.api_archive_file.output_path
}

resource "google_cloudfunctions_function" "api" {
  name    = "satori-codingtest-2-api-04" # FIXME: fix name
  runtime = var.functions_go_runtime

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.zip.name
  trigger_http          = true
  entry_point           = "hello"
  environment_variables = local.envs

  vpc_connector = google_vpc_access_connector.vpc_connector.self_link
}

resource "google_cloudfunctions_function_iam_member" "invoker_api" {
  cloud_function = google_cloudfunctions_function.api.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}

resource "google_cloudfunctions_function" "migration_worker" {
  name    = "migration-worker"
  runtime = var.functions_go_runtime

  timeout               = 540
  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.zip.name
  trigger_http          = true
  entry_point           = "migrate"
  environment_variables = local.envs

  vpc_connector = google_vpc_access_connector.vpc_connector.self_link
}

resource "google_cloudfunctions_function_iam_member" "invoker_migration" {
  cloud_function = google_cloudfunctions_function.migration_worker.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}
