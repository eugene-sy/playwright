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

func (self *Command) SelectFolders() []string {
	result := []string{"tasks"}

	if self.WithHandlers {
		result = append(result, "handlers")
	}
	if self.WithTemplates {
		result = append(result, "templates")
	}
	if self.WithFiles {
		result = append(result, "files")
	}
	if self.WithVars {
		result = append(result, "vars")
	}
	if self.WithDefaults {
		result = append(result, "defaults")
	}
	if self.WithMeta {
		result = append(result, "meta")
	}

	return result
}
