resource "google_artifact_registry_repository" "repo" {
  location      = var.region
  repository_id = "${var.project_id}-web"
  format        = "DOCKER"
}
