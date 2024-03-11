module "network" {
    source = "./network"
    project_id = var.project_id
    network_name = var.network_name
    subnet_name = var.subnet_name
    subnet_region = var.subnet_region
}

module "compute" {
    source = "./compute"
    # global variables
    project_id = var.project_id
   
    # gke variables
    gke_name = var.gke_name
    node_pool_name = var.node_pool_name
    location = module.network.subnet_output.region
    network = module.network.vpc_output.name
    subnetwork = module.network.subnet_output.name
    depends_on = [ module.network ]
}
