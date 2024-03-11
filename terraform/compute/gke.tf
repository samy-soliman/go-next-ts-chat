resource "google_service_account" "gke_service_account" {
  account_id   = "gke-service-account"
  display_name = "gke"
}

resource "google_container_cluster" "gke_primary_cluster" {
    
    name = var.gke_name
    location = var.location
    network = var.network 
    subnetwork = var.subnetwork
    networking_mode = "VPC_NATIVE"
    
    # private_cluster_config {
    #     enable_private_endpoint = false # master
    #     enable_private_nodes = true

    #     # ip address for the control plane
    #     master_ipv4_cidr_block  = "172.16.0.0/28"

    #     # Add the enable-master-global-access flag to create a private cluster
    #     # with global access to the control plane's private endpoint enabled:
    #     master_global_access_config {
    #       enabled = true
    #     }
    # }
    
    ip_allocation_policy {
        cluster_ipv4_cidr_block = "10.11.0.0/21"
        services_ipv4_cidr_block = "10.12.0.0/21"
    }


    master_authorized_networks_config {
    cidr_blocks {
      cidr_block   = "0.0.0.0/0"  # use detailed ip range for security
      display_name = "public_access"
    }
    }
    
    # We can't create a cluster with no node pool defined, but we want to only use
    # separately managed node pools. So we create the smallest possible default
    # node pool and immediately delete it.
    remove_default_node_pool = true
    initial_node_count = 1

    deletion_protection = false

    # enable gatewayAPI API's on the cluster
    gateway_api_config {
      channel = "CHANNEL_STANDARD"
    }
    #depends_on = [google_compute_subnetwork.subnet]
}


resource "google_container_node_pool" "node_pool" {
  name       = var.node_pool_name
  location   = var.location
  cluster    = google_container_cluster.gke_primary_cluster.name
  node_count = 1

  node_config {
    preemptible  = false
    machine_type = "e2-medium"

    // Specify the disk size here
    disk_size_gb = 20

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.gke_service_account.email
    oauth_scopes    = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }

  depends_on = [google_container_cluster.gke_primary_cluster]
}