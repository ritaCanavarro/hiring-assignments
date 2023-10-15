terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.74.0"
    }
  }

  required_version = ">= 0.14"
}

variable "gcp_project" {
  default     = "sre-documentkeeper"
  description = "Project name"
}

variable "gcp_account_id" {
  default     = "sre-documentkeeper-sa"
  description = "Account Id"
}

variable "gcp_region" {
  default     = "europe-central2"
  description = "Region"
}

provider "google" {
  project = var.gcp_project
  region  = var.gcp_region
}

resource "google_project_service" "sre-documentkeeper-compute" {
  project = var.gcp_project
  service = "compute.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}

resource "google_project_service" "sre-documentkeeper-iam" {
  project = var.gcp_project
  service = "iam.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}

resource "google_project_service" "sre-documentkeeper-resourcemanager" {
  project = var.gcp_project
  service = "cloudresourcemanager.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}

resource "google_project_service" "sre-documentkeeper-container" {
  project = var.gcp_project
  service = "container.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}

resource "google_project_service" "sre-documentkeeper-cr" {
  project = var.gcp_project
  service = "containerregistry.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}

resource "google_service_account" "sre-documentkeeper-sa" {
  account_id = var.gcp_account_id
  display_name = "SRE Document Keeper SA"
  project = var.gcp_project
}

resource "google_service_account_key" "sre-documentkeeper-sa_key" {
  service_account_id = var.gcp_account_id
}

resource "google_container_registry" "sre-documentkeeper-registry" {
  project = var.gcp_project
  location = "EU"
}

# VPC
resource "google_compute_network" "vpc" {
  project = var.gcp_project
  name                    = "${var.gcp_project}-vpc"
  auto_create_subnetworks = "false"
}

# Subnet
resource "google_compute_subnetwork" "subnet" {
  project = var.gcp_project
  name          = "${var.gcp_project}-subnet"
  region        = var.gcp_region
  network       = google_compute_network.vpc.name
  ip_cidr_range = "10.10.0.0/24"
}

variable "gke_num_nodes" {
  default     = 3
  description = "number of gke nodes"
}

# GKE cluster
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
}
