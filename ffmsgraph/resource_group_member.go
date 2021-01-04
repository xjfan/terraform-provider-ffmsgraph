package ffmsgraph

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAadGroupMember() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAadGroupMemberCreate,
		ReadContext:   resourceAadGroupMemberRead,
		DeleteContext: resourceAadGroupMemberDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"member_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceAadGroupMemberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	groupID := d.Get("group_id").(string)
	memberID := d.Get("member_id").(string)
	err := c.createAadGroupMember(groupID, memberID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
		return diags
	}

	d.SetId(fmt.Sprintf("%s:%s", groupID, memberID))

	resourceAadGroupMemberRead(ctx, d, m)

	return diags
}

func resourceAadGroupMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	id := d.Id()
	s := strings.Split(id, ":")
	if len(s) != 2 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to parse ID",
		})
		return diags
	}

	groupID := s[0]
	memberID := s[1]

	aadGroupMember, err := c.getAadGroupMember(groupID, memberID)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_id", groupID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("member_id", aadGroupMember.ID); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAadGroupMemberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	groupID := d.Get("group_id").(string)
	memberID := d.Get("member_id").(string)

	err := c.deleteAadGroupMember(groupID, memberID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
