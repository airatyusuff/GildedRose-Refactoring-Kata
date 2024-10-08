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
type SulfurasItem struct{}
type RegularItem struct{}

func (r RegularItem) UpdateItem() {
}

func (ab AgedBrieItem) UpdateItem(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func (bs BackstageItem) UpdateItem(item *Item) {
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
