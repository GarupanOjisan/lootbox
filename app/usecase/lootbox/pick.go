package lootbox

import (
	"fmt"

	model "github.com/garupanojisan/lootbox/app/models/lootbox-item"
	pick "github.com/garupanojisan/rand-pick"
	"github.com/garupanojisan/rand-pick/algorythm"
)

func (u *UseCase) Pick(num int) ([]model.LootBoxItem, error) {
	p := map[int]interface{}{
		70: model.LootBoxItem{Value: "normal"},
		20: model.LootBoxItem{Value: "rare"},
		10: model.LootBoxItem{Value: "super rare"},
	}

	r := pick.RandChoice{
		Probabilities: &p,
		Algorithm:     algorythm.NewWeightedPicker(),
	}
	items, err := r.Pick(num)
	if err != nil {
		return nil, err
	}
	return toLootBoxItems(&items)
}

func toLootBoxItems(items *[]interface{}) ([]model.LootBoxItem, error) {
	lbItems := make([]model.LootBoxItem, len(*items))
	for i, item := range *items {
		switch lbItem := item.(type) {
		case model.LootBoxItem:
			lbItems[i] = lbItem
			break
		default:
			return nil, fmt.Errorf("invalid data")
		}
	}
	return lbItems, nil
}
