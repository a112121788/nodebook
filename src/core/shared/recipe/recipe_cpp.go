package recipe

import (
	"github.com/a112121788/nodebook/src/core/shared/recipe/helper"
	"github.com/a112121788/nodebook/src/core/shared/types"
)

func Cpp() types.Recipe {
	return helper.StdRecipe(
		"cpp",      // key
		"C++14",    // name
		"C++",      // language
		"main.cpp", // mainfile
		"clike",    // cmmode
		"docker.io/library/gcc:latest",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "g++ -std=c++14 -Wall -Wextra -Werror -o /tmp/code.out /code/" + notebook.GetRecipe().GetMainfile() + " && /tmp/code.out"}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "g++ -std=c++14 -Wall -Wextra -Werror -o /tmp/code.out '" + notebook.GetMainFileAbsPath() + "' && /tmp/code.out"}
		},
		nil,
		nil,
	)
}
