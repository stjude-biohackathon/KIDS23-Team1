resource "azurerm_storage_account" "hackathon-hot" {
  name                     = "hackathonhot"
  resource_group_name      = data.azurerm_resource_group.hackathon-rg.name
  location                 = data.azurerm_resource_group.hackathon-rg.location
  
  account_tier             = "Standard"
  account_kind = "StorageV2"
  access_tier              = "Hot"
  account_replication_type = "LRS"
  is_hns_enabled           = true
  public_network_access_enabled = true
  blob_properties {
    delete_retention_policy {
      days = 7
    }
    versioning_enabled = false
    change_feed_enabled = false
  }
  nfsv3_enabled = true
  sftp_enabled = false
  network_rules {
    default_action = "Deny"
  }
  enable_https_traffic_only = true
  shared_access_key_enabled = true
  min_tls_version = "TLS1_0"
  infrastructure_encryption_enabled = false
  large_file_share_enabled = false
  
  tags = {
    environment = "staging"
  }
}

