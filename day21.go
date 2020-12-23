package main

import (
	"fmt"
	"strings"
)

type Composition struct {
	allergensByIngredient []int
	ingredientsByAllergen []int
}

var NO_ALLERGEN int = -2
var UNKNOWN int = -1

func buildComposition(nbrIngredients int, nbrAllergens int) *Composition {
	allergensByIngredient := make([]int, nbrIngredients)
	for i := 0; i < nbrIngredients; i++ {
		allergensByIngredient[i] = UNKNOWN
	}
	ingredientsByAllergen := make([]int, nbrAllergens)
	for i := 0; i < nbrAllergens; i++ {
		ingredientsByAllergen[i] = UNKNOWN
	}
	return &Composition{allergensByIngredient, ingredientsByAllergen}
}

func (comp *Composition) print() {
	// for i := 0; i < level; i++ {
	// 	fmt.Print("  ")
	// }
	fmt.Printf("allergens by Ingredient: %v\n", comp.allergensByIngredient)
	fmt.Printf("ingredients per allergens: %v \n", comp.ingredientsByAllergen)
}

func (comp *Composition) clone() *Composition {
	clonedIngredientsByAllergen := make([]int, len(comp.ingredientsByAllergen))
	copy(clonedIngredientsByAllergen, comp.ingredientsByAllergen)
	clonedAllergenByIngredients := make([]int, len(comp.allergensByIngredient))
	copy(clonedAllergenByIngredients, comp.allergensByIngredient)
	return &Composition{clonedAllergenByIngredients, clonedIngredientsByAllergen}
}

func (comp *Composition) tryAdd(ingredient int, allergen int) bool {
	allrg := comp.allergensByIngredient[ingredient]
	if allrg != UNKNOWN {
		return allrg == allergen
	}

	if allergen != NO_ALLERGEN {
		ingrd := comp.ingredientsByAllergen[allergen]
		if ingrd != UNKNOWN {
			return ingrd == ingredient
		}
	}
	comp.allergensByIngredient[ingredient] = allergen
	if allergen != NO_ALLERGEN {
		comp.ingredientsByAllergen[allergen] = ingredient
	}
	return true
}

func remove(arr []int, element int) []int {
	idx := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			idx = i
			break
		}
	}

	if idx == -1 {
		return arr
	}
	return removeAt(arr, idx)
}

func removeAt(arr []int, index int) []int {
	newArray := make([]int, index)
	copy(newArray, arr[:index])
	if index < len(arr)-1 {
		newArray = append(newArray, arr[index+1:]...)
	}
	return newArray
}

func solveDay21Example() {
	lines := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}
	// lines := getDataFromFile("day21")
	ingredients, nbrIngredients, allergens, nbrAllergens := parseFoodLines(lines)
	for i := 0; i < len(ingredients); i++ {
		fmt.Printf("ingredients: %v\n", ingredients[i])
		fmt.Printf("allergens: %v\n", allergens[i])
	}

	initialSolution := buildComposition(nbrIngredients, nbrAllergens)

	ok, possibleSolution := tryFindPossibleComposition(0, ingredients, allergens, initialSolution)

	if ok {
		possibleSolution.print()
	} else {
		fmt.Printf("Oups, no solution found")
	}
	// ok, solutions := tryFindPossibleCompositionForFood(ingredients[0], allergens[0], solution, 0)
	// fmt.Printf("try to find? %v \n", ok)

	// for i := 1; i < len(ingredients); i++ {
	// 	solutions2 := []*Composition{}
	// 	for _, sol := range solutions {
	// 		ok, newSolutions := tryFindPossibleCompositionForFood(ingredients[1], allergens[1], sol, 0)
	// 		if ok {
	// 			for _, sol2 := range newSolutions {
	// 				solutions2 = append(solutions2, sol2)
	// 			}
	// 		}
	// 	}
	// 	solutions = solutions2
	// }

	// fmt.Printf("*****\n")
	// counts := []int{}
	// for _, sol := range solutions {
	// 	count := 0
	// 	// sol.print(0)
	// 	for _, v := range sol.allergensByIngredient {
	// 		if v == "." {
	// 			count++
	// 		}
	// 	}
	// 	counts = append(counts, count)
	// }
	// fmt.Printf("%v\n", counts)
}

func tryFindPossibleComposition(index int, ingredients [][]int, allergens [][]int, current *Composition) (bool, *Composition) {
	fmt.Printf("index %d \n", index)
	fmt.Printf("ingredients: %v\n", ingredients[index])
	fmt.Printf("allergens: %v\n", allergens[index])
	fmt.Println("current solution")
	current.print()
	fmt.Println("===")
	if index == len(ingredients) {
		return true, current
	}

	foodIngredients := ingredients[index]
	foodAllergens := allergens[index]
	valid, foodIngredients, foodAllergens := current.checkAndTrim(foodIngredients, foodAllergens)
	if !valid {
		fmt.Printf("(idx %d) ingredients and allergens not valid for this solution\n", index)
		return false, nil
	}

	ok, possibleFoodSolutions := tryFindFoodPossibleCompositions(foodIngredients, foodAllergens, current)
	if !ok {
		fmt.Printf("idx %d no solution found\n", index)
		return false, nil
	}

	fmt.Printf("idx %d: %d solutions found\n", index, len(possibleFoodSolutions))

	for _, possibleSolution := range possibleFoodSolutions {
		nextOK, nextSolution := tryFindPossibleComposition(index+1, ingredients, allergens, possibleSolution)
		if nextOK {
			return true, nextSolution
		}
	}
	return false, nil
}

func (comp *Composition) checkAndTrim(ingredients []int, allergens []int) (bool, []int, []int) {
	unusedIngredients := []int{}
	unusedAllergens := make([]int, len(allergens))
	copy(unusedAllergens, allergens)
	for ingredient := range ingredients {
		allrg := comp.allergensByIngredient[ingredient]
		if allrg == UNKNOWN {
			unusedIngredients = append(unusedIngredients, ingredient)
		} else if allrg != NO_ALLERGEN {
			if !contains(allergens, allrg) {
				return false, nil, nil
			} else {
				remove(unusedAllergens, allrg)
			}
		}
	}

	for _, allrg := range unusedAllergens {
		ingr := comp.ingredientsByAllergen[allrg]
		if ingr != UNKNOWN {
			return false, nil, nil
		}
	}

	return true, unusedIngredients, unusedAllergens
}

func tryFindFoodPossibleCompositions(ingredients []int, allergens []int, solution *Composition) (bool, []*Composition) {
	if len(allergens) > len(ingredients) {
		//fmt.Printf("%snot enough ingredients! ==> false\n", prefix)
		return false, nil
	}
	if len(ingredients) == 0 {
		// fmt.Printf("%snot more ingredients! ==> true\n", prefix)
		return true, []*Composition{solution}
	}

	allPossibleSolutions := []*Composition{}

	clonedSolution := solution.clone()
	if clonedSolution.tryAdd(ingredients[0], NO_ALLERGEN) {
		ok, solutions := tryFindFoodPossibleCompositions(ingredients[1:], allergens, clonedSolution)
		if ok {
			allPossibleSolutions = append(allPossibleSolutions, solutions...)
		}
	}
	for i, allrg := range allergens {
		clonedSolution := solution.clone()
		if clonedSolution.tryAdd(ingredients[0], allrg) {
			ok, solutions := tryFindFoodPossibleCompositions(ingredients[1:], removeAt(allergens, i), clonedSolution)
			if ok {
				allPossibleSolutions = append(allPossibleSolutions, solutions...)
			}
		}
	}

	if len(allPossibleSolutions) == 0 {
		return false, nil
	}
	return true, allPossibleSolutions
}

func solveDay21Part1() {
}

// func stringsContain(arr []string, element string) bool {
// 	for _, el := range arr {
// 		if el == element {
// 			return true
// 		}
// 	}
// 	return false
// }

// func tryFindPossibleCompositionForFood(ingredients []int, allergens []int, solution *Composition, level int) (bool, []*Composition) {
// 	//fmt.Printf("%stryFind... for %v and %v\n", prefix, ingredients, allergens)
// 	// solution.print(level)

// 	allPossibleSolutions := []*Composition{}
// 	updatedIngredients := ingredients[1:]
// 	clonedSolution := solution.clone()
// 	if clonedSolution.tryAdd(ingredients[0], ".") {
// 		ok, possibleSolutions := tryFindPossibleCompositionForFood(updatedIngredients, allergens, clonedSolution, level+1)
// 		if ok {
// 			for _, possibleSolution := range possibleSolutions {
// 				allPossibleSolutions = append(allPossibleSolutions, possibleSolution)
// 			}
// 		}
// 	}

// 	for i, allergen := range allergens {
// 		clonedSolution := solution.clone()
// 		if !clonedSolution.tryAdd(ingredients[0], allergen) {
// 			continue
// 		}
// 		updatedAllergens := removeAt(allergens, i)
// 		ok, possibleSolutions := tryFindPossibleCompositionForFood(updatedIngredients, updatedAllergens, clonedSolution, level+1)
// 		if ok {
// 			for _, possibleSolution := range possibleSolutions {
// 				allPossibleSolutions = append(allPossibleSolutions, possibleSolution)
// 			}
// 		}
// 	}

// 	if len(allPossibleSolutions) == 0 {
// 		// fmt.Printf("%scouldn't find any solution! ==> false\n", prefix)

// 		return false, nil
// 	}
// 	// fmt.Printf("%sfound %d possible solutions ==> true\n", prefix, len(allPossibleSolutions))
// 	return true, allPossibleSolutions
// }

func clone(arr [][]string) [][]string {
	arrcopy := make([][]string, len(arr))
	for i := 0; i < len(arr); i++ {
		arrcopy[i] = make([]string, len(arr[i]))
		copy(arrcopy[i], arr[i])
	}
	return arrcopy
}

func solveDay21Part2() {

}

// ----------
func getDay21Data() ([][]int, int, [][]int, int) {
	lines := getDataFromFile("day21")
	return parseFoodLines(lines)
}

func parseFoodLines(lines []string) ([][]int, int, [][]int, int) {
	ingredients := make([][]int, len(lines))
	allergens := make([][]int, len(lines))
	icount, acount := 0, 0
	mapIngredients, mapAllergens := make(map[string]int), make(map[string]int)

	for i, line := range lines {

		idx := strings.Index(line, " (contains")
		for _, singredient := range strings.Split(line[:idx], " ") {
			ingrIdx, ok := mapIngredients[singredient]
			if !ok {
				ingrIdx = icount
				mapIngredients[singredient] = ingrIdx
				icount++
			}
			ingredients[i] = append(ingredients[i], ingrIdx)
		}

		for _, sallergen := range strings.Split(line[idx+11:len(line)-1], ", ") {
			allrgnIdx, ok := mapAllergens[sallergen]
			if !ok {
				allrgnIdx = acount
				mapAllergens[sallergen] = allrgnIdx
				acount++
			}
			allergens[i] = append(allergens[i], allrgnIdx)
		}
	}
	// sort.SliceStable(allergens, func(i, j int) bool { return len(ingredients[i]) < len(ingredients[j]) })
	// sort.SliceStable(ingredients, func(i, j int) bool { return len(ingredients[i]) < len(ingredients[j]) })
	fmt.Printf("There are %d different ingredients and %d allergens\n", len(mapIngredients), len(mapAllergens))
	return ingredients, len(mapIngredients), allergens, len(mapIngredients)
}
