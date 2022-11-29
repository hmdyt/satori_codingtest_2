resource "google_cloud_run_service" "web" {
  name       = "web"
  location   = var.region
  depends_on = [google_service_networking_connection.vpc_connection]

  template {
    spec {
      containers {
        image = var.web_container_uri
      }
    }
  }
}
