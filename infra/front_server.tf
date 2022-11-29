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

  metadata {
    annotations = {
      "autoscaling.knative.dev/maxScale"        = "3"
      "run.googleapis.com/vpc-access-connector" = google_vpc_access_connector.vpc_connector.name
      "run.googleapis.com/vpc-access-egress"    = "all-traffic"
    }
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.web.location
  project  = google_cloud_run_service.web.project
  service  = google_cloud_run_service.web.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
