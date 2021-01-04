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
	aadGroup, err := c.getAadGroupByName(displayName)
	if aadGroup != nil && err == nil {
		d.Set("id", aadGroup.ID)
		d.Set("description", aadGroup.Description)
		d.Set("display_name", aadGroup.DisplayName)
	} else if aadGroup == nil && err == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Can't find this User!",
		})
		return diags
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
		return diags
	}

	d.SetId(aadGroup.ID)

	return diags
}
