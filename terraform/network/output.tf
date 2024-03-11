output "vpc_output" {
  value       =  google_compute_network.vpc   
  description = "network object"
}

output "subnet_output" {
  value       =  google_compute_subnetwork.subnet   
  description = "subnet object"
}