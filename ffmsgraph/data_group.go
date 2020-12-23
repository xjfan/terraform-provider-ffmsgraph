package ffmsgraph

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataAadGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataAadGroupRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataAadGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	displayName := d.Get("display_name").(string)
	aadGroup, _ := c.getAadGroupWithName(displayName)

	if err := d.Set("id", aadGroup.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", aadGroup.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_name", aadGroup.DisplayName); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(aadGroup.ID)

	return diags
}
