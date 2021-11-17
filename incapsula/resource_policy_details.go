package incapsula

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"log"
	"strconv"
)

func resourcePolicyDetails() *schema.Resource {
	return &schema.Resource{
		Create:   resourcepolicy_detailsCreate,
		Read:     resourcepolicy_detailsRead,
		Update:   resourcepolicy_detailsUpdate,
		Delete:   resourcepolicy_detailsDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"policy_id": {
				Description: "The policy ID. During update must be equal to the updated policy ID.",
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				ForceNew:    false,
			},
			"account_id": {
				Description: "Account ID",
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				ForceNew:    false,
			},
			"policy_name": {
				Description: "The name of the policy",
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				ForceNew:    false,
			},
			"status": {
				Description: "Indicates whether policy is enabled or disabled.",
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				ForceNew:    false,
			},
			"sub_category": {
				Description: "Subtype of notification policy. Example values include: ‘account_notifications’; ‘website_notifications’; ‘certificate_management_notifications’",
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				ForceNew:    false,
			},
			"apply_to_new_assets": {
				Description: "If value is ‘TRUE’, all newly onboarded assets are automatically added to the notification policy&#39;s assets list.",
				Type:        schema.TypeBool,
				Optional:    false,
				Required:    true,
				ForceNew:    false,
			},
		},
	}
}

func resourcepolicy_detailsCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("**********[INFO] ****************CREATING RESOURCE")

	apiKeyProviderKey, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-KEY", "2b554ecf-fe9b-4814-beb3-e86ca274dfae")
	apiKeyProviderId, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-ID", "362745")

	demoClient, _ := NewDemoClient("https://api.stage.impervaservices.com/hackathon-notification-service")

	body := CreateJSONRequestBody{
		PolicyId:         d.Get("policy_id").(int),
		AccountId:        d.Get("account_id").(int),
		PolicyName:       d.Get("policy_name").(string),
		Status:           d.Get("status").(string),
		SubCategory:      d.Get("sub_category").(string),
		ApplyToNewAssets: d.Get("apply_to_new_assets").(bool),
	}
	resp, err := demoClient.Create(context.Background(), body, apiKeyProviderKey.Intercept, apiKeyProviderId.Intercept)
	if err != nil {
		return fmt.Errorf("Error from Incapsula service : %s", err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error status code from Incapsula service: %s", string(responseBody))
	}

	// Parse the JSON
	var response ImpervaApiResponseNotificationPolicyFull
	err = json.Unmarshal([]byte(responseBody), &response)
	var notificationPolicyFull = response.Data
	d.SetId(strconv.Itoa(notificationPolicyFull.PolicyId))
	return resourcepolicy_detailsRead(d, m)
}

func resourcepolicy_detailsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("**********[INFO] ****************READING RESOURCE")

	apiKeyProviderKey, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-KEY", "2b554ecf-fe9b-4814-beb3-e86ca274dfae")
	apiKeyProviderId, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-ID", "362745")

	demoClient, _ := NewDemoClient("https://api.stage.impervaservices.com/hackathon-notification-service")

	policyId, err := strconv.Atoi(d.Id())

	resp, err := demoClient.Get(context.Background(), policyId, apiKeyProviderKey.Intercept, apiKeyProviderId.Intercept)
	if err != nil {
		return fmt.Errorf("Error from Incapsula service : %s", err)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error status code from Incapsula service: %s", string(responseBody))
	}

	// Parse the JSON
	var response ImpervaApiResponseNotificationPolicyFull
	var notificationPolicyFull *NotificationPolicyFull
	log.Printf("read :")
	log.Printf("%v", string(responseBody))
	log.Printf("%v", response)
	log.Printf("%v", notificationPolicyFull)
	err = json.Unmarshal([]byte(responseBody), &response)
	notificationPolicyFull = response.Data
	d.Set("policy_id", notificationPolicyFull.PolicyId)
	d.Set("account_id", notificationPolicyFull.AccountId)
	d.Set("policy_name", notificationPolicyFull.PolicyName)
	d.Set("status", notificationPolicyFull.Status)
	d.Set("sub_category", notificationPolicyFull.SubCategory)
	d.Set("apply_to_new_assets", notificationPolicyFull.ApplyToNewAssets)

	return nil
}

func resourcepolicy_detailsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("**********[INFO] ****************UPDATING RESOURCE")
	apiKeyProviderKey, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-KEY", "2b554ecf-fe9b-4814-beb3-e86ca274dfae")
	apiKeyProviderId, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-ID", "362745")

	demoClient, _ := NewDemoClient("https://api.stage.impervaservices.com/hackathon-notification-service")

	body := UpdateJSONRequestBody{
		PolicyId:         d.Get("policy_id").(int),
		AccountId:        d.Get("account_id").(int),
		PolicyName:       d.Get("policy_name").(string),
		Status:           d.Get("status").(string),
		SubCategory:      d.Get("sub_category").(string),
		ApplyToNewAssets: d.Get("apply_to_new_assets").(bool),
	}

	resourceId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	resp, err := demoClient.Update(context.Background(), resourceId, body, apiKeyProviderKey.Intercept, apiKeyProviderId.Intercept)
	if err != nil {
		return fmt.Errorf("Error from Incapsula service : %s", err)
	}
	log.Printf("update done :&v", resp)

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error status code from Incapsula service: %s", string(responseBody))
	}

	var response ImpervaApiResponseNotificationPolicyFull
	var notificationPolicyFull *NotificationPolicyFull
	log.Printf("create :")
	log.Printf("%v", string(responseBody))
	log.Printf("%v", response)
	log.Printf("%v", notificationPolicyFull)
	err = json.Unmarshal([]byte(responseBody), &response)
	notificationPolicyFull = response.Data
	d.Set("policy_id", notificationPolicyFull.PolicyId)
	d.Set("account_id", notificationPolicyFull.AccountId)
	d.Set("policy_name", notificationPolicyFull.PolicyName)
	d.Set("status", notificationPolicyFull.Status)
	d.Set("sub_category", notificationPolicyFull.SubCategory)
	d.Set("apply_to_new_assets", notificationPolicyFull.ApplyToNewAssets)
	return nil
}

func resourcepolicy_detailsDelete(d *schema.ResourceData, m interface{}) error {
	//demoClient := m.(*Client)
	//delete method code!
	return nil
}
