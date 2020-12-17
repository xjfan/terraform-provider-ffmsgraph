package ffmsgraph

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// func DataAadGroup() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: dataAadGroupRead,

// 		Schema: map[string]*schema.Schema{
// 			"value": &schema.Schema{
// 				Type:     schema.TypeList,
// 				Computed: true,
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"id": {
// 							Type:     schema.TypeString,
// 							Optional: true,
// 							Computed: true,
// 						},
// 						"description": {
// 							Type:     schema.TypeString,
// 							Optional: true,
// 						},
// 						"display_name": {
// 							Type:     schema.TypeString,
// 							Required: true,
// 							ForceNew: true,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

func DataAadGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataAadGroupRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func dataAadGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	display_name := d.Get("display_name").(string)
	ings, _ := c.getAadGroup(display_name)

	log.Printf("[#] ings:", ings)
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
