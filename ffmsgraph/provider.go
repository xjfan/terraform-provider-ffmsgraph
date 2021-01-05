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
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"FFMSGRAPH_AZURE_OBJECT_ID", "ARM_OBJECT_ID"}, nil),
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"FFMSGRAPH_AZURE_TENANT_ID", "ARM_TENANT_ID"}, nil),
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"FFMSGRAPH_AZURE_CLIENT_ID", "ARM_CLIENT_ID"}, nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"FFMSGRAPH_AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET"}, nil),
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
