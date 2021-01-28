package ffmsgraph

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAadInvitedUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAadInvitedUserCreate,
		ReadContext:   resourceAadInvitedUserRead,
		DeleteContext: resourceAadInvitedUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			// `id` in invitedUser
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"invited_user_email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"invite_redirect_url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAadInvitedUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	invitedUserEmailAddress := d.Get("invited_user_email").(string)
	inviteRedirectURL := d.Get("invite_redirect_url").(string)

	aadInvitedUser, err := c.postAadInvitedUser(invitedUserEmailAddress, inviteRedirectURL)
	if aadInvitedUser == nil && err ！= nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
		return diags
	d.SetId(aadInvitedUser.ID)
	resourceAadInvitedUserRead(ctx, d, m)
	return nil
}

func resourceAadInvitedUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	aadInvitedUserID := d.Id()

	aadGroup, err := c.getAadGroup(aadInvitedUserID)
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
