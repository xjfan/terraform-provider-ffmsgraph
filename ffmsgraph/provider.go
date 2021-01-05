package ffmsgraph

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"ffmsgraph_group": DataAadGroup(),
			"ffmsgraph_user":  DataAadUser(),
			"ffmsgraph_app":   DataAadApp(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ffmsgraph_group":        ResourceAadGroup(),
			"ffmsgraph_group_member": ResourceAadGroupMember(),
		},
		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ConfigureContextFunc: configureProvider,
	}
	return p
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	objectID := d.Get("object_id").(string)
	tenantID := d.Get("tenant_id").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	client, err := APIClient(objectID, tenantID, clientID, clientSecret)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create oauth client",
		})
		return nil, diags
	}

	return client, diags
}
