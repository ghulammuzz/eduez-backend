provider "google" {
#   credentials = file("")
  project     = "shhhh-gcp-project-id"
  region      = "asia-southeast2"
}

resource "google_compute_instance" "eduze_instance" {
  name         = "eduze-instance"
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
    access_config {
      
    }
  }

  tags = ["web", "dev"]
}
