package ffmsgraph

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAadGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAadGroupCreate,
		ReadContext:   resourceAadGroupRead,
		DeleteContext: resourceAadGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
				ForceNew: true,
			},
		},
	}
}

func resourceAadGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	displayName := d.Get("display_name").(string)

	aadGroup, err := c.getAadGroupByName(displayName)
	if aadGroup != nil && err == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "This AadGroup existed in AAD!",
		})
		return diags
	} else if aadGroup == nil && err == nil {
		aadGroup, err := c.createAadGroup(displayName)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
			return diags
		}
		d.SetId(aadGroup.ID)
		resourceAadGroupRead(ctx, d, m)
		return diags
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
		return diags
	}
}

func resourceAadGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	aadGroupID := d.Id()

	aadGroup, err := c.getAadGroup(aadGroupID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
		return diags
	}
	d.Set("id", aadGroup.ID)
	d.Set("description", aadGroup.Description)
	d.Set("display_name", aadGroup.DisplayName)
	return diags
}

func resourceAadGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	aadGroupID := d.Id()

	err := c.deleteAadGroup(aadGroupID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
		return diags
	}

	d.SetId("")

	return diags
}
