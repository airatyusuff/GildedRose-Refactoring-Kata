package gildedrose

import (
	"strings"
)

const MAX_SELLIN_FOR_DOUBLE_PRICE_BACKSTAGE = 10
const MAX_SELLIN_FOR_TRIPLE_PRICE_BACKSTAGE = 5
const MAX_ITEM_QUALITY = 50
const MIN_ITEM_QUALITY = 0
const TYPE_CONJURED = "conjured"
const TYPE_AGED_BRIE = "aged brie"
const TYPE_SULFURAS = "sulfuras"
const TYPE_BACKSTAGE = "backstage passes"

func isItemQualityValid(quality int) bool {
	return quality > MIN_ITEM_QUALITY && quality < MAX_ITEM_QUALITY
}

func UpdateSellInDate(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}
}

func UpdateQuality(items []*Item) {
	categorisedInventoryItems := map[string]InventoryItem{}

	for _, item := range items {
		categoriseInventoryItem(item.Name, categorisedInventoryItems)
	}
	UpdateInventoryItems(items, categorisedInventoryItems)
}

func UpdateInventoryItems(items []*Item, categorisedItems map[string]InventoryItem) {
	for _, item := range items {
		categorisedItem := categorisedItems[item.Name]
		categorisedItem.UpdateItem(item)
	}
}

func categoriseInventoryItem(itemName string, inventoryItems map[string]InventoryItem) {
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
