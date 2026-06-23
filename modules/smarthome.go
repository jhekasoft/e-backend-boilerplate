//go:build smarthome || all
// +build smarthome all

package modules

import "e-backend-boilerplate/modules/smarthome"

func init() {
	m := smarthome.NewModule()
	EnabledModules = append(EnabledModules, m)
}
