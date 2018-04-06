---
layout: "cf"
page_title: "Cloud Foundry: cf_space_quota"
sidebar_current: "docs-cf-resource-space-quota"
description: |-
  Provides a Cloud Foundry space Quota resource.
---

# cf\_space\_quota

Provides a Cloud Foundry resource to manage org or space [quotas](https://docs.cloudfoundry.org/adminguide/quota-plans.html) definitions.

~> **NOTE:** This resource requires the provider to be authenticated with an account granted admin permissions.


## Example Usage

The following example creates a quota within an Org referenced by `cf_org.myorg.id` that can be applied to any space within that Org.

```
resource "cf_space_quota" "10g" {
    name = "10g"
    allow_paid_service_plans = false
    instance_memory = 512
    total_memory = 10240
    total_app_instances = 10
    total_routes = 5
    total_services = 20
    org = "${cf_org.myorg.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name you use to identify the quota or plan in Cloud Foundry
* `allow_paid_service_plans` - (Required) Determines whether users can provision instances of non-free service plans. Does not control plan visibility. When false, non-free service plans may be visible in the marketplace but instances can not be provisioned.
* `instance_memory` - (Optional) Maximum memory per application instance
* `total_memory` - (Required) Maximum memory usage allowed
* `total_app_instances` - (Optional) Maximum app instances allowed
* `total_routes` - (Required) Maximum routes allowed
* `total_services` - (Required) Maximum services allowed
* `total_route_ports` - (Optional) Maximum routes with reserved ports
* `total_private_domains` - (Optional) Maximum number of private domains allowed to be created within the Org
* `org` - (Optional) The Org within which this quota will be defined so it can be applied to spaces within that Org

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the quota

## Import

The current space Quota can be imported using the `quota`, e.g.

```
$ terraform import cf_space_quota.10g a-guid
```
