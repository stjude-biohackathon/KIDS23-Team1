resource "azurerm_storage_account" "hackathon-cold" {
  name                     = "hackathoncold"
  resource_group_name      = data.azurerm_resource_group.hackathon-rg.name
  location                 = data.azurerm_resource_group.hackathon-rg.location
  
  account_tier             = "Standard"
  account_kind = "StorageV2"
  access_tier              = "Cool"
  account_replication_type = "GRS"
  is_hns_enabled           = false
  public_network_access_enabled = false
  blob_properties {
    delete_retention_policy {
      days = 7
    }
    container_delete_retention_policy {
      days = 7
    }
    versioning_enabled = true
    change_feed_enabled = true
  }
  nfsv3_enabled = false
  cross_tenant_replication_enabled = true
  enable_https_traffic_only = true
  shared_access_key_enabled = true
  min_tls_version = "TLS1_0"
  infrastructure_encryption_enabled = false
  network_rules {
    default_action = "Deny"
  }
  routing {
    choice = "MicrosoftRouting"
  }
  
  tags = {
    environment = "staging"
  }
}

