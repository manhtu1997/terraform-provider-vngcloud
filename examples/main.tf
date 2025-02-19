terraform {
  required_providers {
    vngcloud = {
      source  = "vngcloud/vngcloud"
      version = "0.0.14"
    }
  }
}

provider "vngcloud" {
  token_url     = "https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token"
  client_id     = var.client_id
  client_secret = var.client_secret
  vserver_base_url      = "https://vserverapi.vngcloud.vn/vserver-gateway"
}

data "vngcloud_vserver_volume_type_zone" "volume_type_zone" {
  name       = "SSD"
  project_id = var.project_id
}
data "vngcloud_vserver_volume_type" "volume_type" {
  name                = var.volume_type_name
  project_id          = var.project_id
  volume_type_zone_id = data.vngcloud_vserver_volume_type_zone.volume_type_zone.id
}

resource "vngcloud_vserver_server" "server" {
  count             = var.server_count
  project_id        = var.project_id
  name              = "vngcloud-server-${count.index}"
  encryption_volume = false
  attach_floating   = true
  flavor_id         = var.flavor_id
  image_id          = var.image_id
  network_id        = var.network_id
  root_disk_size    = var.root_disk_size
  root_disk_type_id = data.vngcloud_vserver_volume_type.volume_type.id
  ssh_key           = var.ssh_key_id
  security_group    = var.security_group_id_list
  subnet_id         = var.subnet_id
  action            = "start"
  user_name         = "stackops"
  user_password     = "Vng@Cloud3030"
  lifecycle {
    create_before_destroy = true
  }
}

resource "vngcloud_vserver_volume" "volume" {
  count          = var.server_count
  name           = "vngcloud-volume-${count.index}"
  size           = var.data_disk_size
  volume_type_id = data.vngcloud_vserver_volume_type.volume_type.id
  project_id     = var.project_id
  lifecycle {
    create_before_destroy = true
  }
}

resource "vngcloud_vserver_volume_attach" "attach_volume" {
  count       = var.server_count
  project_id  = var.project_id
  volume_id   = vngcloud_vserver_volume.volume[count.index].id
  server_id = vngcloud_vserver_server.server[count.index].id
}

data "vngcloud_vserver_server_group_policy" "policy" {
  name = var.server_group_policy_name
  project_id = var.project_id
}

resource "vngcloud_vserver_server_group" "server_group" {
  description = "description of server group"
  name        = "example-server-group"
  policy_id   = data.vngcloud_vserver_server_group_policy.policy.id
  project_id = var.project_id
}
