variable "username" {
  description = "The username for the DB master user"
  type        = string
  sensitive   = true
}
variable "password" {
  description = "The password for the DB master user"
  type        = string
  sensitive   = true
}

terraform {
  required_providers {
    vkcs = {
      source = "vk-cs/vkcs"
    }
  }
}

provider "vkcs" {
  # Your user account.
  username = var.username

  # The password of the account
  password = var.password

  # The tenant token can be taken from the project Settings tab - > API keys.
  # Project ID will be our token.
  project_id = "8ede2ee5cc98414fa34d9df66e236530"

  # Region name
  region = "RegionOne"
}