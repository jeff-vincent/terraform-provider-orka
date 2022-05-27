terraform {
  required_providers {
    orka = {
      source = "macstadium.com/edu/orka"
    }
  }
}

provider "orka" {
  host = "http://10.221.188.100"
  email = "email@email.com"
  password = "password"
  license_key = "<YOUR_KEY>"
}