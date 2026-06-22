package models

type Module interface {
	Name() string
	Run(core *Core) error
}
