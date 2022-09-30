package recipe

import (
	"github.com/a112121788/nodebook/src/core/shared/service"
)

func AddRecipesToRegistry(recipeRegistry *service.RecipeRegistry) {

	recipeRegistry.
		AddRecipe(C()).
		AddRecipe(Clojure()).
		AddRecipe(Cpp()).
		AddRecipe(Csharp()).
		AddRecipe(Elixir()).
		AddRecipe(Fsharp()).
		AddRecipe(Go()).
		AddRecipe(Haskell()).
		AddRecipe(Java()).
		AddRecipe(Lua()).
		AddRecipe(NodeJS()).
		AddRecipe(Ocaml()).
		AddRecipe(Php()).
		AddRecipe(Python3()).
		AddRecipe(R()).
		AddRecipe(Ruby()).
		AddRecipe(Rust()).
		AddRecipe(Swift()).
		AddRecipe(Typescript())
}
