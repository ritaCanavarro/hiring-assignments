# GKE cluster
variable "gke_num_nodes" {
  default     = 3
  description = "number of gke nodes"
}

data "google_container_engine_versions" "gke_version" {
  project = var.gcp_project
  location = var.gcp_region
  version_prefix = "1.27."
}

resource "google_container_cluster" "primary" {
  project = var.gcp_project
  name     = "${var.gcp_project}-gke"
  location = var.gcp_region

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  network    = google_compute_network.vpc.name
  subnetwork = google_compute_subnetwork.subnet.name
}

# Separately Managed Node Pool
resource "google_container_node_pool" "primary_nodes" {
  project    =    var.gcp_project
  name       = google_container_cluster.primary.name
  location   = var.gcp_region
  cluster    = google_container_cluster.primary.name
  
  version = data.google_container_engine_versions.gke_version.release_channel_latest_version["STABLE"]
  node_count = var.gke_num_nodes

  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]

    labels = {
      env = var.gcp_project
    }

    disk_size_gb = 50
    machine_type = "n1-standard-1"
    tags         = ["gke-node", "${var.gcp_project}-gke"]
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }

  timeouts {
    create = "40m"
    update = "50m"
  }
}
