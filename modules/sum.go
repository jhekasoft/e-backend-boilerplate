//go:build sum || all
// +build sum all

package modules

import "e-backend-boilerplate/modules/sum"

func init() {
	m := sum.NewModule()
	EnabledModules = append(EnabledModules, m)
}
