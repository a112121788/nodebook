package recipe

import (
	"github.com/a112121788/nodebook/src/core/shared/recipe/helper"
	"github.com/a112121788/nodebook/src/core/shared/types"
)

func Java() types.Recipe {
	return helper.StdRecipe(
		"java",      // key
		"Java",      // name
		"Java",      // language
		"main.java", // mainfile
		"clike",     // cmmode
		"docker.io/library/java:latest",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", `javac -d /tmp "/code/` + notebook.GetRecipe().GetMainfile() + `" && cd /tmp && java Main`}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", `javac -d /tmp "` + notebook.GetMainFileAbsPath() + `" && cd /tmp && java Main`}
		},
		nil,
		nil,
	)
}
