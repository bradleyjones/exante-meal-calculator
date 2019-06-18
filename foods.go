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
	// Proteins
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
	// Vegetables
	"Alfalfa Sporouts":        &Food{23, 0.2, 0},
	"Asparagus":               &Food{20, 1.7, 0},
	"Aubergine":               &Food{25, 3.0, 0},
	"Avocado":                 &Food{160, 2.0, 0},
	"Bamboo Shoots":           &Food{27, 2.8, 0},
	"Bean Sprouts":            &Food{30, 4, 0},
	"Brocoli":                 &Food{34, 4.4, 0},
	"Cabbage (any type)":      &Food{25, 3.5, 0},
	"Cauliflower":             &Food{25, 3.0, 0},
	"Celery":                  &Food{16, 1.4, 0},
	"Courgette":               &Food{17, 2.1, 0},
	"Cucumber":                &Food{16, 3.1, 0},
	"Fennel":                  &Food{31, 3.9, 0},
	"Green Beans":             &Food{31, 4.3, 0},
	"Green Pepper":            &Food{20, 2.9, 0},
	"Lettuce":                 &Food{15, 1.6, 0},
	"Kale":                    &Food{28, 3.6, 0},
	"Marrow":                  &Food{11, 1.1, 0},
	"Mushrooms":               &Food{25, 3.0, 0},
	"Olives (green or black)": &Food{146, 0.4, 0},
	"Pak Choi":                &Food{13, 1.2, 0},
	"Radish":                  &Food{16, 2.0, 0},
	"Rocket":                  &Food{25, 2.0, 0},
	"Spinach":                 &Food{23, 1.4, 0},
	"Spring Onions":           &Food{32, 4.4, 0},
	"Tomato":                  &Food{18, 3.0, 0},
	"Turnip":                  &Food{22, 3.0, 0},
	"Watercress":              &Food{11, 0.8, 0},
}
