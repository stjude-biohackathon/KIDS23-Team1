terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_virtual_network" "hackathon-vnet" {
  name                = "hackathon-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = data.azurerm_resource_group.hackathon-rg.location
  resource_group_name = data.azurerm_resource_group.hackathon-rg.name
}

resource "azurerm_network_security_group" "hackathon-nsg" {
  name = "hackathon-nsg"
  location = data.azurerm_resource_group.hackathon-rg.location
  resource_group_name = data.azurerm_resource_group.hackathon-rg.name
  security_rule {
    name = "sj-ssh-in"
    priority = 101
    direction = "Inbound"
    access = "Allow"
    protocol = "Tcp"
    source_port_range = "*"
    destination_port_range = "22"
    source_address_prefix = "192.55.208.0/24"
    destination_address_prefix = "VirtualNetwork"
    description = "allow ssh from SJ"
  }
}

resource "azurerm_subnet" "hackathon-subnet" {
  name                 = "hackathon-subnet"
  resource_group_name  = data.azurerm_resource_group.hackathon-rg.name
  virtual_network_name = azurerm_virtual_network.hackathon-vnet.name
  address_prefixes     = ["10.0.2.0/24"]
  service_endpoints = ["Microsoft.Storage"]
}

resource "azurerm_public_ip" "hackathon-pip" {
  name                = "hackathonpip"
  resource_group_name = data.azurerm_resource_group.hackathon-rg.name
  location            = data.azurerm_resource_group.hackathon-rg.location
  allocation_method   = "Static"
}

resource "azurerm_network_interface" "hackathon-nic" {
  name                = "hackathon-nic"
  location            = data.azurerm_resource_group.hackathon-rg.location
  resource_group_name = data.azurerm_resource_group.hackathon-rg.name
  depends_on = [azurerm_public_ip.hackathon-pip]

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.hackathon-subnet.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id = azurerm_public_ip.hackathon-pip.id
  } 
}

resource "azurerm_linux_virtual_machine" "hackathon-vm" {
  name                = "hackathon-vm"
  resource_group_name = data.azurerm_resource_group.hackathon-rg.name
  location            = data.azurerm_resource_group.hackathon-rg.location
  size                = "Standard_D8s_v5"
  admin_username      = "dkennetz"
  network_interface_ids = [
    azurerm_network_interface.hackathon-nic.id,
  ]

  admin_ssh_key {
    username   = "dkennetz"
    public_key = file("~/.ssh/id_rsa.pub")
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
    name = "hackathon-vm-osdisk"
    disk_size_gb = 200
  }

  source_image_reference {
    publisher = "OpenLogic"
    offer     = "CentOS-HPC"
    sku       = "7.6"
    version   = "7.6.2020062900"
  }
}

