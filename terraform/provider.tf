provider "google" {
    # credentials = file("SA_KEY.json")
    project = var.project_id
    region =  var.project_region   # default for resources
}