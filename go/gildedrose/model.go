package gildedrose

type BackstageItem struct {
	item *Item
}
type AgedBrieItem struct {
	item *Item
}
type ConjuredItem struct {
	item *Item
}
type SulfurasItem struct {
	item *Item
}
type RegularItem struct {
	item *Item
}

func (s SulfurasItem) UpdateItem() {}

func (r RegularItem) UpdateItem() {
	UpdateSellInDate(r.item)

	if isItemPastSellIn(r.item.SellIn) {
		r.item.Quality = r.item.Quality - 2
		return
	}

	if r.item.Quality > MIN_ITEM_QUALITY {
		r.item.Quality = r.item.Quality - 1
	}
}

func (c ConjuredItem) UpdateItem() {
	UpdateSellInDate(c.item)
	if c.item.Quality >= MIN_ITEM_CONJURED_QUALITY {
		c.item.Quality = c.item.Quality - 2
	}
}

func (ab AgedBrieItem) UpdateItem() {
	UpdateSellInDate(ab.item)
	if ab.item.Quality < MAX_ITEM_QUALITY {
		ab.item.Quality = ab.item.Quality + 1
	}
}

func (bs BackstageItem) UpdateItem() {
	UpdateSellInDate(bs.item)
	if bs.item.SellIn < 0 {
		bs.item.Quality = 0
		return
	}

	if isItemQualityValid(bs.item.Quality) && isBackstageItemForTripleIncrease(bs.item.SellIn) {
		bs.item.Quality = bs.item.Quality + 3
		return
	}

	if isItemQualityValid(bs.item.Quality) && isBackstageItemForDoubleIncrease(bs.item.SellIn) {
		bs.item.Quality = bs.item.Quality + 2
		return
	}

	if isItemQualityValid(bs.item.Quality) && isBackstageItemForRegularIncrease(bs.item.SellIn) {
		bs.item.Quality = bs.item.Quality + 1
		return
	}
}

func UpdateSellInDate(item *Item) {
	item.SellIn = item.SellIn - 1
}

func isItemPastSellIn(sellIn int) bool {
	return sellIn < 0
}

func isItemQualityValid(quality int) bool {
	return quality >= MIN_ITEM_QUALITY && quality < MAX_ITEM_QUALITY
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
