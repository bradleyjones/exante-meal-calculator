package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func addToMealList(meal_list *tview.List, key string, increase bool) {
	search := meal_list.FindItems(key, "", true, true)
	if len(search) == 0 && increase {
		var calories float32 = float32(foods[key].Calories) / float32(100)
		foods[key].Grams = 1
		meal_list.AddItem(key+" - 1g", fmt.Sprintf("%.1f Calories", calories), 0, nil)
	} else if len(search) == 1 {
		quantity := foods[key].Grams
		if increase {
			quantity++
		} else {
			quantity--
		}
		if quantity == 0 {
			foods[key].Grams = 0
			meal_list.RemoveItem(search[0])
		} else if quantity >= 1 {
			foods[key].Grams = quantity
			var calories float32 = (float32(foods[key].Calories) / float32(100)) * float32(quantity)
			meal_list.SetItemText(search[0], key+fmt.Sprintf(" - %dg", quantity), fmt.Sprintf("%.1f Calories", calories))
		}
	}
}

func updateMealDescription(meal_list *tview.List, description *tview.TextView) {
	var totalCalories float32 = 0.0
	for i := 0; i < meal_list.GetItemCount(); i++ {
		itemText, _ := meal_list.GetItemText(i)
		item := strings.Split(itemText, " -")[0]
		totalCalories += (float32(foods[item].Calories) / float32(100)) * float32(foods[item].Grams)
	}
	description.SetText(fmt.Sprintf("Total Calories for meal: %.1f", totalCalories))
}

func main() {
	app := tview.NewApplication()

	meal_list := tview.NewList()
	meal_description := tview.NewTextView()
	food_list := tview.NewList()
	food_list.SetSelectedFunc(func(i int, k, v string, x rune) {
		addToMealList(meal_list, k, true)
		updateMealDescription(meal_list, meal_description)
	})

	for _, k := range sortedKeys(foods) {
		description := "Calories (per 100g): " + fmt.Sprintf("%d", foods[k].Calories) +
			" Carbs (per 100g): " + fmt.Sprintf("%.2f", foods[k].Carbs)
		food_list.AddItem(k, description, 0, nil)
	}

	main_grid := tview.NewGrid().
		SetSize(2, 2, 0, 0).
		SetBorders(true).
		AddItem(food_list, 0, 0, 2, 1, 0, 0, false).
		AddItem(meal_list, 0, 1, 1, 1, 0, 0, false).
		AddItem(meal_description, 1, 1, 1, 1, 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			app.SetFocus(meal_list)
		} else if event.Key() == tcell.KeyCtrlP {
			app.SetFocus(food_list)
		} else if event.Key() == tcell.KeyBackspace2 {
			k, _ := food_list.GetItemText(food_list.GetCurrentItem())
			addToMealList(meal_list, k, false)
			updateMealDescription(meal_list, meal_description)
		}
		return event
	})
	if err := app.SetRoot(main_grid, true).SetFocus(food_list).Run(); err != nil {
		panic(err)
	}
}
