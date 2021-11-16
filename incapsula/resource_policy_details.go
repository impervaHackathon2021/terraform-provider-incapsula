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
			"policyid": {
				Description: "The policy ID. During update must be equal to the updated policy ID.",
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				ForceNew:    false,
			},
			"accountid": {
				Description: "Account ID",
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				ForceNew:    false,
			},
			"policyname": {
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
			"subcategory": {
				Description: "Subtype of notification policy. Example values include: ‘account_notifications’; ‘website_notifications’; ‘certificate_management_notifications’",
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				ForceNew:    false,
			},
			"applytonewassets": {
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

	apiKeyProviderKey, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-KEY", "2b554ecf-fe9b-4814-beb3-e86ca274dfae")
	apiKeyProviderId, _ := securityprovider.NewSecurityProviderApiKey("header", "x-API-ID", "362745")

	demoClient, _ := NewDemoClient("https://api.stage.impervaservices.com/hackathon-notification-service")

	body := CreateJSONRequestBody{
		PolicyId:         d.Get("policyid").(int),
		AccountId:        d.Get("accountid").(int),
		PolicyName:       d.Get("policyname").(string),
		Status:           d.Get("status").(string),
		SubCategory:      d.Get("subcategory").(string),
		ApplyToNewAssets: d.Get("applytonewassets").(bool),
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
	log.Printf("response create unmarshaled:")
	log.Printf("%v", notificationPolicyFull)
	d.SetId(strconv.Itoa(notificationPolicyFull.PolicyId))
	return resourcepolicy_detailsRead(d, m)
}

func resourcepolicy_detailsRead(d *schema.ResourceData, m interface{}) error {

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
	err = json.Unmarshal([]byte(responseBody), &response)
	var notificationPolicyFull = response.Data

	d.Set("policyid", notificationPolicyFull.PolicyId)
	d.Set("accountid", notificationPolicyFull.AccountId)
	d.Set("policyname", notificationPolicyFull.PolicyName)
	d.Set("status", notificationPolicyFull.Status)
	d.Set("subcategory", notificationPolicyFull.SubCategory)
	d.Set("applytonewassets", notificationPolicyFull.ApplyToNewAssets)

	return nil
}

func resourcepolicy_detailsUpdate(d *schema.ResourceData, m interface{}) error {
	//demoClient := m.(*Client)
	//update  method code!
	return resourcepolicy_detailsRead(d, m)
}

func resourcepolicy_detailsDelete(d *schema.ResourceData, m interface{}) error {
	//demoClient := m.(*Client)
	//delete method code!
	return nil
}
