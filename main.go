package main

type Ingredient struct {
	Name     string
	Quantity string
}

type Recipe struct {
	Name        string
	Ingredients []Ingredient
}

func worker(recipes <-chan Recipe, done chan<- int) {
	for recipe := range recipes {
		done <- len(recipe.Ingredients)
	}
}

func main() {

}
