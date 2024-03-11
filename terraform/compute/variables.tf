# project id  
variable "project_id" {
    type = string    
}

##### gke variables
variable "gke_name" {
    type = string      
}
variable "node_pool_name" {
    type = string      
}
variable "location" {
    type = string       # module.network.subnet_out.region
}
variable "network" {
    type = string       # module.network.vpc_out.name
}
variable "subnetwork" {
    type = string       # module.network.subnet_out.name
}


