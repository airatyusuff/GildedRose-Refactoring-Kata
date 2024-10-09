package gildedrose

const MAX_SELLIN_FOR_DOUBLE_PRICE_BACKSTAGE = 10
const MAX_SELLIN_FOR_TRIPLE_PRICE_BACKSTAGE = 5
const MAX_ITEM_QUALITY = 50
const SULFURAS_ITEM_QUALITY = 80
const MIN_ITEM_QUALITY = 0
const MIN_ITEM_CONJURED_QUALITY = 2

type Item struct {
	Name            string
	SellIn, Quality int
}

type InventoryItem interface {
	UpdateItem()
}

func UpdateInventory(items []*Item) {
	for _, item := range items {
		inventoryItem := createInventoryItemByName(item)
		inventoryItem.UpdateItem()
	}
}

func createInventoryItemByName(item *Item) InventoryItem {
	switch item.Name {
	case "Backstage passes to a TAFKAL80ETC concert":
		return BackstageItem{item: item}
	case "Aged Brie":
		return AgedBrieItem{item: item}
	case "Sulfuras, Hand of Ragnaros":
		return SulfurasItem{item: item}
	case "Conjured Mana Cake":
		return ConjuredItem{item: item}
	default:
		return RegularItem{item: item}
	}
}
