terraform {
  required_providers {
    orka = {
      source = "macstadium.com/edu/orka"
    }
  }
}

provider "orka" {
  email = "email@email.com"
  password = "password"
  license_key = "<YOUR_KEY>"
}