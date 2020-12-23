package main

import (
	"fmt"
	"strings"
)

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

// ----------
func getDay21Data() ([][]int, int, [][]int, int, map[string]int, map[string]int) {
	lines := getDataFromFile("day21")
	return parseFoodLines(lines)
}

func parseFoodLines(lines []string) ([][]int, int, [][]int, int, map[string]int, map[string]int) {
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
	fmt.Printf("There are %d different ingredients and %d allergens\n", len(mapIngredients), len(mapAllergens))
	return ingredients, len(mapIngredients), allergens, len(mapAllergens), mapIngredients, mapAllergens
}

func intersect(arr1 []int, arr2 []int) []int {
	inter := []int{}
	for _, el := range arr1 {
		if contains(arr2, el) {
			inter = append(inter, el)
		}
	}
	return inter
}

func solveDay21Example() {
	lines := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}
	ingredients, nbrIngredients, allergens, nbrAllergens, mapIngredients, mapAllergens := parseFoodLines(lines)

	ingredientsPerAllergen := solve(ingredients, allergens, nbrAllergens)
	allergenNames := make([]string, nbrAllergens)
	for k, v := range mapAllergens {
		allergenNames[v] = k
	}
	ingredientsNames := make([]string, nbrIngredients)
	for k, v := range mapIngredients {
		ingredientsNames[v] = k
	}
	dangerousList := make([]string, nbrAllergens)
	for i, ingr := range ingredientsPerAllergen {
		dangerousList[i] = ingredientsNames[ingr]
	}
	fmt.Printf("allergen: %v\n", allergenNames)
	fmt.Printf("dangerous list: %v\n", dangerousList)
}

func solveDay21Part1() {
	lines := getDataFromFile("day21")
	ingredients, _, allergens, nbrAllergens, _, _ := parseFoodLines(lines)

	solve(ingredients, allergens, nbrAllergens)
}

func solveDay21Part2() {
	lines := getDataFromFile("day21")

	ingredients, nbrIngredients, allergens, nbrAllergens, mapIngredients, mapAllergens := parseFoodLines(lines)

	ingredientsPerAllergen := solve(ingredients, allergens, nbrAllergens)
	allergenNames := make([]string, nbrAllergens)
	for k, v := range mapAllergens {
		allergenNames[v] = k
	}
	ingredientsNames := make([]string, nbrIngredients)
	for k, v := range mapIngredients {
		ingredientsNames[v] = k
	}
	dangerousList := make([]string, nbrAllergens)
	for i, ingr := range ingredientsPerAllergen {
		dangerousList[i] = ingredientsNames[ingr]
	}
	fmt.Printf("allergen: %v\n", allergenNames)
	fmt.Printf("dangerous list: %v\n", dangerousList)
}

func solve(ingredients [][]int, allergens [][]int, nbrAllergens int) []int {
	possibleIngredientsPerAllergens := make([][]int, nbrAllergens)
	for index := 0; index < len(ingredients); index++ {
		for _, allergen := range allergens[index] {
			if possibleIngredientsPerAllergens[allergen] == nil {
				clone := ingredients[index]
				possibleIngredientsPerAllergens[allergen] = clone
			} else {
				possibleIngredientsPerAllergens[allergen] = intersect(possibleIngredientsPerAllergens[allergen], ingredients[index])
			}
		}
	}

	used := make([]bool, nbrAllergens)
	for true {
		idx := -1
		for i, possibleIngredients := range possibleIngredientsPerAllergens {
			if len(possibleIngredients) == 1 && !used[i] {
				idx = i
				used[idx] = true
				break
			}
		}
		if idx == -1 {
			break
		}
		ingr := possibleIngredientsPerAllergens[idx][0]
		for i := 0; i < len(possibleIngredientsPerAllergens); i++ {
			if i != idx {
				possibleIngredientsPerAllergens[i] = remove(possibleIngredientsPerAllergens[i], ingr)
			}
		}
	}

	ingredientsPerAllergen := make([]int, nbrAllergens)
	for i, possibleIngredients := range possibleIngredientsPerAllergens {
		// fmt.Printf("%v\n", possibleIngredients)
		if len(possibleIngredients) > 1 {
			fmt.Printf("Too many ingredients (%d) for allergen %d \n", len(possibleIngredients), i)
			return nil
		}
		ingredientsPerAllergen[i] = possibleIngredients[0]
	}

	fmt.Printf("Ingredients per allergens: %v\n", ingredientsPerAllergen)

	count := 0
	for i := 0; i < len(ingredients); i++ {
		for j := 0; j < len(ingredients[i]); j++ {
			if !contains(ingredientsPerAllergen, ingredients[i][j]) {
				count++
			}
		}
	}
	fmt.Printf("Ingredients with no allergen appear %d times\n", count)
	return ingredientsPerAllergen
}

// allergen: [shellfish  fish    peanuts  nuts    eggs  dairy    wheat   soy]
// dangerous:[bqtvr      bmrmhm  vflms    snhrpv  zmb   bqkndvb  rkkrx   qzkjrtl]

// sorted allergens: dairy eggs fish nuts peanuts shellfish soy wheat
// sorted dangerous: bqkndvb,zmb,bmrmhm,snhrpv,vflms,bqtvr,qzkjrtl,rkkrx
