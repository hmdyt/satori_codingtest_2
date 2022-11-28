resource "google_sql_database_instance" "sql" {
  name                = "sql"
  database_version    = "MYSQL_8_0"
  deletion_protection = true
  depends_on          = [google_service_networking_connection.vpc_connection]
  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.vpc.id
    }
  }
}
