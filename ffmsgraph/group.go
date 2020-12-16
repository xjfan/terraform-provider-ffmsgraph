package ffmsgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
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
	c := m.(Client)

	displayName := d.Get("display_name").(string)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:7000/groups/%s", displayName), nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create AadGroup client",
		})

		return diags
	}

	r, err := client.Do(req)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get AadGroup client request",
		})

		return diags
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	if r.StatusCode == 200 {
		diags = append(diags, diag.Diagnostic{
			Summary: fmt.Sprintf(c.Token),
			Detail:  fmt.Sprintf(string(body)),
		})
		return diags
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
