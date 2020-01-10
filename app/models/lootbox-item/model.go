package lootbox_item

import "time"

type LootBoxItem struct {
	ID          string
	GroupID     string
	Value       string
	Probability int
	createdAt   time.Time
	updatedAt   time.Time
}
