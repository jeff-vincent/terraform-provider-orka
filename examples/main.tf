terraform {
  required_providers {
    orka = {
      source = "macstadium.com/edu/orka"
    }
  }
}

provider "orka" {
<<<<<<< HEAD
  host = "http://10.221.188.100"
  email = "email@email.com"
  password = "password"
  license_key = "<YOUR_KEY>"
}
=======

}

data "orka_vms" "vms" {

}
>>>>>>> 453c57f96b4152d6e10bb893285e7f0923d2b425
