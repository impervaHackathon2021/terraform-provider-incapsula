---
layout: "incapsula"
page_title: "Incapsula: incap-policy-details"
sidebar_current: "docs-incapsula-resource-policy-details"
description: |- Provides a Incapsula Policy Details resource.
---

# incapsula_policy_details

Provides an Incapsula Policy Details resource.

Policy Details Imperva Hackathon-Notofications API Documentation.

## Example Usage

```hcl
resource "incapsula_policy_details" "demo-terraform-policy-details" {
     policy_id = 
     account_id = 1234
     policy_name = new policy
     status = ENABLE
     sub_category = NETWORK_CONNECTIVITY
     apply_to_new_assets = false
}
```

## Argument Reference

The following arguments are supported:

* `policy_id` - (Optional) The policy ID. During update must be equal to the updated policy ID..
* `account_id` - (Optional) Account ID.
* `policy_name` - (Required) The name of the policy.
* `status` - (Required) Indicates whether policy is enabled or disabled..
* `sub_category` - (Required) Subtype of notification policy. Example values include: ‘account_notifications’; ‘website_notifications’; ‘certificate_management_notifications’.
* `apply_to_new_assets` - (Required) If value is ‘TRUE’, all newly onboarded assets are automatically added to the notification policy&#39;s assets list..

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the API Security Site Configuration.
