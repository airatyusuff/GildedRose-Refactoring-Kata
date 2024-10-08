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

// func isSulfurItem(item *Item) bool {
// 	return item.Name == "Sulfuras, Hand of Ragnaros"
// }

// func isRegularItem(itemName string) bool {
// 	itemNameLower := strings.ToLower(itemName)
// 	return (itemName != "Aged Brie" &&
// 		itemName != "Sulfuras, Hand of Ragnaros" &&
// 		!strings.Contains(itemNameLower, "conjured") &&
// 		!strings.Contains(itemNameLower, "backstage passes"))
// }

func isItemQualityValid(quality int) bool {
	return quality > MIN_ITEM_QUALITY && quality < MAX_ITEM_QUALITY
}

// func isItemMaxQuality(quality int) bool {
// 	return quality == MAX_ITEM_QUALITY
// }

// func updateBackstageItem(item *Item) {
// 	if item.SellIn < MIN_ITEM_QUALITY {
// 		item.Quality = MIN_ITEM_QUALITY
// 		return
// 	}
// 	if isItemQualityValid(item.Quality) && isItemInRangeForTripleIncrease(item.SellIn) {
// 		item.Quality = item.Quality + 3
// 		return
// 	}
// 	if isItemQualityValid(item.Quality) && isItemForDoubleIncrease(item.SellIn) {
// 		item.Quality = item.Quality + 2
// 		return
// 	}
// 	item.Quality = item.Quality + 1
// }

// func updateConjuredItem(item *Item) {
// 	if isItemQualityValid(item.Quality) {
// 		item.Quality = item.Quality - 2
// 	}
// }

// func updateAgedBrieItem(item *Item) {
// 	if item.Quality < MAX_ITEM_QUALITY {
// 		item.Quality = item.Quality + 1
// 	}
// }

// func updateRegularItem(item *Item) {
// 	item.Quality = item.Quality - 1
// }

func UpdateSellInDate(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}
}

// func UpdateItemQuality(item *Item) {
// 	if isSulfurItem(item) {
// 		return
// 	}
// 	updateSellInDate(item)
// 	if isRegularItem(item.Name) && isItemQualityValid(item.Quality) {
// 		updateRegularItem(item)
// 		return
// 	}
// 	if item.Quality < MAX_ITEM_QUALITY {
// 		itemNameLower := strings.ToLower(item.Name)
// 		if strings.Contains(itemNameLower, "conjured") {
// 			updateConjuredItem(item)
// 			return
// 		}
// 		if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
// 			updateBackstageItem(item)
// 			return
// 		}
// 		if item.Name == "Aged Brie" {
// 			updateAgedBrieItem(item)
// 			return
// 		}
// 		item.Quality = item.Quality + 1
// 	}
// }

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
