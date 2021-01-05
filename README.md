# Terraform Provider for Microsoft Graph

## Introduction

This Terraform Provider is the plugin to provision and manage AAD resources via GraphAPI to match the requirements for Farfetch Infra Data Platform

You need Terraform v0.13 and an Azure AD tenant with the following privilege:
- Directory.Read.All
- Group.Create
- Group.Read.All
- User.Read

## Purpose

The Infra Data Platform needs to use AAD to manage the authentication on ACL management in ADLS Gen2

## Supported resources

- Data sources
  - ffmsgraph_group
  - ffmsgraph_user
  - ffmsgraph_app
- Resources
  - ffmsgraph_group
  - ffmsgraph_group_member

## Provider configuration

The credential for provider should be managed by Data Infra Team

```hcl
provider "ffmsgraph" {
  object_id           = "" // env: FFMSGRAPH_AZURE_OBJECT_ID, ARM_OBJECT_ID
  tenant_id           = "" // env: FFMSGRAPH_AZURE_TENANT_ID, ARM_TENANT_ID
  client_id           = "" // env: FFMSGRAPH_AZURE_CLIENT_ID, ARM_CLIENT_ID
  client_secret       = "" // env: FFMSGRAPH_AZURE_CLIENT_SECRET, ARM_CLIENT_SECRET
}
```

## Tips
The following features and limitations need to be known before implemetation:
- Don't support to create duplicated name for AAD group (Actually it allows officially via GraphAPI)
- Each Creation AAD group by this provider will bind provider tenant as owner for management
- The schema of resources haven't totally followed the schema of Azure Microsoft Graph, just use what we need

## How to test

Terraform v0.13 and Go v1.14 are required.

```console
$ git clone git@github.com:farfetch-internal/terraform-provider-ffmsgraph.git
$ go build -o ~/.terraform.d/plugins/data-infra/xjfan/ffmsgraph/1.0.0/darwin_amd64/terraform-provider-ffmsgraph # MacOS
```

```console
terraform {
  required_providers {
    msgraph = {
      source = "hashicorp.com/xjfan/ffmsgraph"
      version = "1.0"
    }
  }
}
```

Run terraform with an environment variable `TF_LOG=DEBUG` to enable debug log output:

```console
$ terraform init
$ TF_LOG=DEBUG terraform plan
$ TF_LOG=DEBUG terraform apply
```
Import resource with objectID
```console
terraform import ffmsgraph_group_member.member1 8df64495-bd29-49c7-b895-42c1c07d760a:722a8eef-c32a-4316-b928-89fc78a9ee7e
terraform import ffmsgraph_group.aadgroup1 8df64495-bd29-49c7-b895-42c1c07d760a
...
```

## Template

```console
data "ffmsgraph_group" "aadgroup" {
  display_name = "adls_test_ro"
}

data "ffmsgraph_user" "aaduser" {
  mail = "jim.fan@farfetch.com"
}

data "ffmsgraph_app" "aadapp" {
  app_id = "b1a358fa-938b-49d9-8b65-a5cd468f220d"
}

output "aadgroup_display_name" {
  value = data.ffmsgraph_group.aadgroup.display_name
}

output "aaduser_display_name" {
  value = data.ffmsgraph_user.aaduser.display_name
}

output "aadapp_display_name" {
  value = data.ffmsgraph_app.aadapp.display_name
}

resource "ffmsgraph_group" "aadgroup1" {
  display_name = "adls_test1_ro"
}

resource "ffmsgraph_group_member" "member1" {
  group_id = ffmsgraph_group.aadgroup1.id
  member_id = data.ffmsgraph_user.aaduser.id
}

resource "ffmsgraph_group_member" "member2" {
  group_id = ffmsgraph_group.aadgroup1.id
  member_id = data.ffmsgraph_group.aadgroup.id
}

resource "ffmsgraph_group_member" "member3" {
  group_id = ffmsgraph_group.aadgroup1.id
  member_id = data.ffmsgraph_app.aadapp.id
}
```



## Todo

- [ ] To add test module
- [ ] Work on integration with other terraform provider
- [ ] OAuth2 tokens in backend storage
- [ ] Publish to Terraform registry
- [ ] CI/CD (GoReleaser)
