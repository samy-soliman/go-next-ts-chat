# VPC network is a global entity spanning all GCP regions
resource "google_compute_network" "vpc" {
    name = var.network_name
    auto_create_subnetworks = false
    project = var.project_id
}