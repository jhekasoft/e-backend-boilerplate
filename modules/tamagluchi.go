//go:build tamagluchi || all
// +build tamagluchi all

package modules

import "e-backend-boilerplate/modules/tamagluchi"

func init() {
	m := tamagluchi.NewModule()
	EnabledModules = append(EnabledModules, m)
}
