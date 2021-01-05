package ffmsgraph

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataAadApp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataAadAppRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataAadAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	appID := d.Get("app_id").(string)
	aadApp, err := c.getAadApp(appID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
		return diags
	}
	d.Set("id", aadApp.ID)
	d.Set("app_id", aadApp.AppID)
	d.Set("display_name", aadApp.DisplayName)

	d.SetId(aadApp.ID)

	return diags
}
