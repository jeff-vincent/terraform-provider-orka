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

output "vm_resources" {
  value = data.orka_vms.vms
}


resource "orka_vms" "edu" {
  vms {
    orka_vm_name = "myorkavm1"
	  orka_base_image = "bigsur-ssh-git.img"
	  orka_cpu_core = 3
	  vcpu_count = 3
  }
}

output "edu_vms" {
  value = orka_vms.edu
}