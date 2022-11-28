resource "google_compute_network" "vpc" {
  name                    = "vpc"
  auto_create_subnetworks = false
}

resource "google_compute_global_address" "private_id_address" {
  name          = "private-ip-address"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.vpc.id
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.vpc.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_id_address.name]
}

resource "google_vpc_access_connector" "vpc_connector" {
  name          = "vpc-connector"
  ip_cidr_range = "10.14.0.0/28"
  network       = google_compute_network.vpc.name
}
