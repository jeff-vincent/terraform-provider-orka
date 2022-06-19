terraform {
  required_providers {
    orka = {
      source = "macstadium.com/edu/orka"
    }
  }
}

provider "orka" {

}

# data "orka_vms" "vms" {

# }

# output "vm_resources" {
#   value = data.orka_vms.vms
# }


resource "vm_configs" "edu" {
  vm_config {
    orka_vm_name = "myorkavm"
	  orka_base_image = "bigsur-ssh-git.img"
	  orka_cpu_core = 3
	  vcpu_count = 3
  }
  vm_config {
    orka_vm_name = "myorkavm1"
	  orka_base_image = "bigsur-ssh-git.img"
	  orka_cpu_core = 3
	  vcpu_count = 3
  }
  vm_config {
    orka_vm_name = "myorkavm11"
	  orka_base_image = "bigsur-ssh-git.img"
	  orka_cpu_core = 3
	  vcpu_count = 3
  }
}

output "vm_configs" {
  value = orka_vm_configs.edu
}

resource "deploy_vm_configs" {

}

