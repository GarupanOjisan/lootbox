package lootbox

import model "github.com/garupanojisan/omikuji-api/app/models/lootbox-item-group"

func (u *UseCase) GetGroup(userId string, id string) (model.LootBoxItemGroup, error) {
	return model.LootBoxItemGroup{
		ID: id,
	}, nil
}
