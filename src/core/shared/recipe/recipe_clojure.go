package recipe

import (
	"github.com/a112121788/nodebook/src/core/shared/recipe/helper"
	"github.com/a112121788/nodebook/src/core/shared/types"
)

func Clojure() types.Recipe {
	return helper.StdRecipe(
		"clojure",   // key
		"Clojure",   // name
		"Clojure",   // language
		"index.clj", // mainfile
		"clojure",   // cmmode
		"docker.io/library/clojure:tools-deps",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "clojure " + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "clojure '" + notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
	)
}
