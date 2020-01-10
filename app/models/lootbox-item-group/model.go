package lootbox_item_group

import "time"

type LootBoxItemGroup struct {
	ID        string
	UserID    string
	Title     string
	createdAt time.Time
	updatedAt time.Time
}
