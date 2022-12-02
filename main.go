// SPDX-FileCopyrightText: Magenta ApS
// SPDX-License-Identifier: MPL-2.0
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-os2mo/os2mo"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return os2mo.Provider()
		},
	})
}
