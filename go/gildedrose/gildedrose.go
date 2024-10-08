package gildedrose

import (
	"strings"
)

const MAX_SELLIN_FOR_DOUBLE_PRICE_BACKSTAGE = 10
const MAX_SELLIN_FOR_TRIPLE_PRICE_BACKSTAGE = 5
const MAX_ITEM_QUALITY = 50
const MIN_ITEM_QUALITY = 0
const MIN_ITEM_CONJURED_QUALITY = 2
const TYPE_CONJURED = "conjured"
const TYPE_AGED_BRIE = "aged brie"
const TYPE_SULFURAS = "sulfuras"
const TYPE_BACKSTAGE = "backstage passes"

type Item struct {
	Name            string
	SellIn, Quality int
}

type InventoryItem interface {
	UpdateItem(item *Item)
}

func UpdateInventory(items []*Item) {
	inventoryItems := map[string]InventoryItem{}

	for _, item := range items {
		categoriseInventoryItemByName(item.Name, inventoryItems)
	}

	updateInventoryItems(items, inventoryItems)
}

func categoriseInventoryItemByName(itemName string, inventoryItems map[string]InventoryItem) {
	itemNameLower := strings.ToLower(itemName)

	if strings.Contains(itemNameLower, TYPE_AGED_BRIE) {
		inventoryItems[itemName] = AgedBrieItem{}
		return
	}

	if strings.Contains(itemNameLower, TYPE_BACKSTAGE) {
		inventoryItems[itemName] = BackstageItem{}
		return
	}

	if strings.Contains(itemNameLower, TYPE_CONJURED) {
		inventoryItems[itemName] = ConjuredItem{}
		return
	}

	if strings.Contains(itemNameLower, TYPE_SULFURAS) {
		inventoryItems[itemName] = SulfurasItem{}
		return
	}

	inventoryItems[itemName] = RegularItem{}
}

func updateInventoryItems(items []*Item, categorisedItems map[string]InventoryItem) {
	for _, item := range items {
		categorisedItem := categorisedItems[item.Name]
		categorisedItem.UpdateItem(item)
	}
}
