resource "google_compute_subnetwork" "subnet" {
  name          = var.subnet_name
  ip_cidr_range = "10.1.0.0/24"
  region        = var.subnet_region
  network       = google_compute_network.vpc.id
  secondary_ip_range {
    range_name    = "gke-secondary-range"
    ip_cidr_range = "10.2.0.0/24"
  }
  depends_on = [google_compute_network.vpc]
}