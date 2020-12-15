package ffmsgraph

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"ffaad_group": DataAadGroup(),
		},
		ResourcesMap: map[string]*schema.Resource{},
		Schema:       map[string]*schema.Schema{
			"token": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
    },
		ConfigureFunc: configureProvider,
	addContextToAllResources(p)
	return p
}

func configureProvider(data *schema.ResourceData) (interface{}, error) {
  opts := &ffmsgraphOpts {
    token:      data.Get("token").(string),
  }

  client, err := APIClient(opts)

  return client, nil
}