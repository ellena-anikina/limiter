package main

import (
	"reflect"
	"slices"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	recipes := []Recipe{
		{
			Name: "Pizza",
			Ingredients: []Ingredient{
				{"Flour", "2 cups"},
				{"Yeast", "1 packet"},
				{"Water", "1 cup"},
			},
		},
		{
			Name: "Pasta",
			Ingredients: []Ingredient{
				{"Pasta", "200g"},
				{"Water", "2 liters"},
				{"Salt", "1 teaspoon"},
				{"Olive oil", "1 tablespoon"},
			},
		},
	}

	recipeChannel := make(chan Recipe, len(recipes))
	done := make(chan int, len(recipes))

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(recipeChannel, done)
		}()
	}

	go func() {
		for _, recipe := range recipes {
			recipeChannel <- recipe
		}
		close(recipeChannel)
	}()

	wg.Wait()
	close(done)

	expected := []int{3, 4}
	results := make([]int, 0, len(recipes))
	for result := range done {
		results = append(results, result)
	}

	slices.Sort(results)

	if !reflect.DeepEqual(expected, results) {
		t.Errorf("Expected %v, got %v", expected, results)
	}
}
