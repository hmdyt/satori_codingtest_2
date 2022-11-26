terraform {
  backend "gcs" {
    bucket = "satori-codingtest-2-tfstete"
    prefix = "terraform/state"
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}
