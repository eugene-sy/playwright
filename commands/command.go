package commands

type Command struct {
	PlaybookName string
	WithHandlers bool
	WithTemplates bool
	WithFiles bool
	WithVars bool
	WithDefaults bool
	WithMeta bool
}

type ICommand interface {
	Execute() (err error)
}
