# project id
variable "project_id" {
  type = string
}

# project default region
variable "project_region" {
  type = string
}

# network module variables
variable "network_name" {
  type = string
}

variable "subnet_name" {
  type = string
}

variable "subnet_region" {
  type = string
}

# compute module variables
variable "gke_name" {
  type = string
}

variable "node_pool_name" {
  type = string
}