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
	inventoryItems := map[string]InventoryItem{}

	for _, item := range items {
		addItemToInventoryByName(item, inventoryItems)
	}

	updateInventoryItems(items, inventoryItems)
}

func addItemToInventoryByName(item *Item, inventoryItems map[string]InventoryItem) {
	switch item.Name {
	case "Backstage passes to a TAFKAL80ETC concert":
		sa := &BackstageItem{item: item}
		inventoryItems[item.Name] = sa
	case "Aged Brie":
		sa := &AgedBrieItem{item: item}
		inventoryItems[item.Name] = sa
	case "Sulfuras, Hand of Ragnaros":
		sa := &SulfurasItem{item: item}
		inventoryItems[item.Name] = sa
	case "Conjured Mana Cake":
		sa := &ConjuredItem{item: item}
		inventoryItems[item.Name] = sa
	default:
		sa := &RegularItem{item: item}
		inventoryItems[item.Name] = sa
	}
}

func updateInventoryItems(items []*Item, categorisedItems map[string]InventoryItem) {
	for _, item := range items {
		categorisedItem := categorisedItems[item.Name]
		categorisedItem.UpdateItem()
	}
}
