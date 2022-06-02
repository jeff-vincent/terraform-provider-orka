terraform {
  required_providers {
    orka = {
      source = "macstadium.com/edu/orka"
    }
  }
}

provider "orka" {

}

data "orka_vms" "vms" {

}

output "vm_stuff" {
  value = data.orka_vms.vms
}