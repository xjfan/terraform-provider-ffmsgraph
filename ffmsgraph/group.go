package ffmsgraph

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataAadGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataAadGroupRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"display": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataAadGroupRead(data *schema.ResourceData, meta interface{}) error {
	opts := &AadGroupOpts {
	  configuration:   {"display": data.Get("display").(string)}
	  endpoint_suffix: "/groups"
	}
  
	responseMap, err := APIClient(opts)

	if _, exists := responseMap["id"]; exists {
	  data.Set("id", responseMap["id"])
	} else {
	  fmt.Errorf("id not found in response from AadGroup.")
	}
	return err
  }