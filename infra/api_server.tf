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
  runtime = "go116"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.zip.name
  trigger_http          = true
  entry_point           = "hello"

}

resource "google_cloudfunctions_function_iam_member" "invoker" {
  cloud_function = google_cloudfunctions_function.api.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}

