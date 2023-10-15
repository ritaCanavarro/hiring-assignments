terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.74.0"
    }
  }

  required_version = ">= 0.14"
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
