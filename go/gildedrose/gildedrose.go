package gildedrose

import (
	"strconv"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func isRegularItem(itemName string) bool {
	return (itemName != "Aged Brie" &&
		itemName != "Backstage passes to a TAFKAL80ETC concert" &&
		itemName != "Sulfuras, Hand of Ragnaros")
}

func isItemQualityValid(quality int) bool {
	return quality > 0 && quality < 50
}

func updateBackstageItem(item *Item) {
	if item.SellIn < 11 && isItemQualityValid(item.Quality) {
		item.Quality = item.Quality + 1
	}
	if item.SellIn < 6 && isItemQualityValid(item.Quality) {
		item.Quality = item.Quality + 1
	}
}

func updateRegularItem(item *Item) {
	item.Quality = item.Quality - 1
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		UpdateItemQuality(item)
	}
}

func updateSellInDate(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}
}

func UpdateItemQuality(item *Item) {
	updateSellInDate(item)

	if isRegularItem(item.Name) && isItemQualityValid(item.Quality) {
		updateRegularItem(item)
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1

			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				updateBackstageItem(item)
			}
		}
	}

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if item.Quality > 0 {
					if item.Name != "Sulfuras, Hand of Ragnaros" {
						item.Quality = item.Quality - 1
					}
				}
			} else {
				item.Quality = 0
			}
		} else {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
}

func OutputInventory(items []*Item) []string {
	days := 2
	output := make([]string, 0)

	for day := 0; day < days; day++ {
		for _, item := range items {
			itemAsStr := item.Name + " " + strconv.Itoa(item.SellIn) + " " + strconv.Itoa(item.Quality)
			output = append(output, itemAsStr)
		}
		UpdateQuality(items)
	}
	return output
}
