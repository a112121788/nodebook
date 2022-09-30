package recipe

import (
	"github.com/a112121788/nodebook/src/core/shared/recipe/helper"
	"github.com/a112121788/nodebook/src/core/shared/types"
)

func Go() types.Recipe {
	return helper.StdRecipe(
		"go",      // key
		"Go",      // name
		"Go",      // language
		"main.go", // mainfile
		"go",      // cmmode
		"docker.io/library/golang:latest",
		func(notebook types.Notebook) []string {
			return []string{"go", "run", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"go", "run", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
	)
}
