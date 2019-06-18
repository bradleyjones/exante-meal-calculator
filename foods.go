package main

import "sort"

type Food struct {
	Calories int32   // Per 100grams
	Carbs    float32 // Per 100grams
	Grams    int32
}

func sortedKeys(m map[string]*Food) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

var foods = map[string]*Food{
	"Cottage Cheese":          &Food{98, 3.4, 0},
	"Quorn Pieces":            &Food{86, 4, 0},
	"Tofu":                    &Food{76, 1.6, 0},
	"Walnuts":                 &Food{610, 4.4, 0},
	"Eggs (2 Medium is 100g)": &Food{150, 1, 0},
	"Tuna":                    &Food{116, 0, 0},
	"Kippers":                 &Food{217, 0, 0},
	"Salmon":                  &Food{208, 0, 0},
	"Prawns":                  &Food{105, 0, 0},
	"Sea bass/white fish":     &Food{125, 0.5, 0},
	"Chicken breast":          &Food{165, 0, 0},
	"Turkey breast":           &Food{104, 0, 0},
	"Bacon":                   &Food{292, 0, 0},
	"Liver":                   &Food{167, 1, 0},
	"Ground beef":             &Food{136, 0, 0},
	"Lamb mince":              &Food{215, 0, 0},
	"Pork Chops":              &Food{231, 0, 0},
}
