package gildedrose

import "strings"

func isSulfurItem(item *Item) bool {
	return item.Name == "Sulfuras, Hand of Ragnaros"
}

func isRegularItem(itemName string) bool {
	itemNameLower := strings.ToLower(itemName)

	return (itemName != "Aged Brie" &&
		itemName != "Sulfuras, Hand of Ragnaros" &&
		!strings.Contains(itemNameLower, "conjured") &&
		!strings.Contains(itemNameLower, "backstage passes"))
}

func isItemQualityValid(quality int) bool {
	return quality > 0 && quality < 50
}

func updateBackstageItem(item *Item) {
	if item.SellIn < 0 {
		item.Quality = 0
		return
	}
	if isItemQualityValid(item.Quality) {
		if item.SellIn < 6 {
			item.Quality = item.Quality + 3
			return
		}
		if item.SellIn < 11 {
			item.Quality = item.Quality + 2
			return
		}
		item.Quality = item.Quality + 1
	}
}

func updateConjuredItem(item *Item) {
	if isItemQualityValid(item.Quality) {
		item.Quality = item.Quality - 2
	}
}

func updateAgedBrieItem(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func updateRegularItem(item *Item) {
	item.Quality = item.Quality - 1
}

func updateSellInDate(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}
}

func UpdateItemQuality(item *Item) {
	if isSulfurItem(item) {
		return
	}

	updateSellInDate(item)

	if isRegularItem(item.Name) && isItemQualityValid(item.Quality) {
		updateRegularItem(item)
		return
	}

	if item.Quality < 50 {
		itemNameLower := strings.ToLower(item.Name)

		if strings.Contains(itemNameLower, "conjured") {
			updateConjuredItem(item)
			return
		}

		if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			updateBackstageItem(item)
			return
		}
		if item.Name == "Aged Brie" {
			updateAgedBrieItem(item)
			return
		}

		item.Quality = item.Quality + 1
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		UpdateItemQuality(item)
	}
}
