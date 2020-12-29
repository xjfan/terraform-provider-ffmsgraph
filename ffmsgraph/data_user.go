package ffmsgraph

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataAadUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataAadUserRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mail": {
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

func dataAadUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	mail := d.Get("mail").(string)
	aadUser, _ := c.getAadGroupByMail(mail)

	if err := d.Set("id", aadUser.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mail", aadUser.Mail); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_name", aadUser.DisplayName); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(aadUser.ID)

	return diags
}
