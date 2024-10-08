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
	switch itemName {
	case "Backstage passes to a TAFKAL80ETC concert":
		inventoryItems[itemName] = BackstageItem{}
	case "Aged Brie":
		inventoryItems[itemName] = AgedBrieItem{}
	case "Sulfuras, Hand of Ragnaros":
		inventoryItems[itemName] = SulfurasItem{}
	case "Conjured Mana Cake":
		inventoryItems[itemName] = ConjuredItem{}
	default:
		inventoryItems[itemName] = RegularItem{}
	}
}

func updateInventoryItems(items []*Item, categorisedItems map[string]InventoryItem) {
	for _, item := range items {
		categorisedItem := categorisedItems[item.Name]
		categorisedItem.UpdateItem(item)
	}
}
