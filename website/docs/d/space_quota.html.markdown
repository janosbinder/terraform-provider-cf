---
layout: "cf"
page_title: "Cloud Foundry: cf_space_quota"
sidebar_current: "docs-cf-datasource-space-quota"
description: |-
  Get information on a Cloud Foundry space Quota.
---

# cf\_space\_quota

Gets information on a Cloud Foundry space quota.

## Example Usage

The following example looks up a space quota named 'myquota' within the Org identified by the id of an Org resource defined elsewhere in the Terraform configuration. 

```
data "cf_space_quota" "q" {
    name = "myquota"
    org = "${cf_org.o1.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the space quota to look up
* `org` - (Required) The organization within which the quota is defined

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the space quota
