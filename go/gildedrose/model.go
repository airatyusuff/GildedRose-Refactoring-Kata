package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type InventoryItem interface {
	UpdateItem(item *Item)
}

type BackstageItem struct{}
type AgedBrieItem struct{}
type ConjuredItem struct{}
type SulfurasItem struct{}
type RegularItem struct{}

func (s SulfurasItem) UpdateItem(item *Item) {}

func (r RegularItem) UpdateItem(item *Item) {
	UpdateSellInDate(item)
	item.Quality = item.Quality - 1
}

func (ab AgedBrieItem) UpdateItem(item *Item) {
	UpdateSellInDate(item)
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func (c ConjuredItem) UpdateItem(item *Item) {
	UpdateSellInDate(item)
	if isItemQualityValid(item.Quality) {
		item.Quality = item.Quality - 2
	}
}

func (bs BackstageItem) UpdateItem(item *Item) {
	UpdateSellInDate(item)
	if item.SellIn < 0 {
		item.Quality = 0
		return
	}

	if isItemQualityValid(item.Quality) && isBackstageItemForTripleIncrease(item.SellIn) {
		item.Quality = item.Quality + 3
		return
	}

	if isItemQualityValid(item.Quality) && isBackstageItemForDoubleIncrease(item.SellIn) {
		item.Quality = item.Quality + 2
		return
	}

	if isItemQualityValid(item.Quality) && isBackstageItemForRegularIncrease(item.SellIn) {
		item.Quality = item.Quality + 1
		return
	}
}

func isBackstageItemForTripleIncrease(sellIn int) bool {
	return sellIn < MAX_SELLIN_FOR_TRIPLE_PRICE_BACKSTAGE
}

func isBackstageItemForDoubleIncrease(sellIn int) bool {
	return sellIn > 0 && sellIn < MAX_SELLIN_FOR_DOUBLE_PRICE_BACKSTAGE
}

func isBackstageItemForRegularIncrease(sellIn int) bool {
	return sellIn > MAX_SELLIN_FOR_DOUBLE_PRICE_BACKSTAGE
}
