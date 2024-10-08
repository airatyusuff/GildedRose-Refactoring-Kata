package gildedrose_test

import (
	"strconv"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestInventoryUpdate(t *testing.T) {
	items := []*gildedrose.Item{
		{"Conjured sample item", 0, 13},
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 20},
		{"Conjured Mana Cake", 3, 6},
	}

	output := OutputInventory(items)

	expected := []string{
		"Conjured sample item 0 13",
		"+5 Dexterity Vest 10 20",
		"Aged Brie 2 0",
		"Elixir of the Mongoose 5 0",
		"Sulfuras, Hand of Ragnaros 0 80",
		"Sulfuras, Hand of Ragnaros -1 80",
		"Backstage passes to a TAFKAL80ETC concert 15 20",
		"Backstage passes to a TAFKAL80ETC concert 10 50",
		"Backstage passes to a TAFKAL80ETC concert 0 20",
		"Conjured Mana Cake 3 6",
		"Conjured sample item -1 11",
		"+5 Dexterity Vest 9 19",
		"Aged Brie 1 1",
		"Elixir of the Mongoose 4 0",
		"Sulfuras, Hand of Ragnaros 0 80",
		"Sulfuras, Hand of Ragnaros -1 80",
		"Backstage passes to a TAFKAL80ETC concert 14 21",
		"Backstage passes to a TAFKAL80ETC concert 9 50",
		"Backstage passes to a TAFKAL80ETC concert -1 0",
		"Conjured Mana Cake 2 4",
	}

	for i, tc := range expected {
		if tc != output[i] {
			t.Errorf("Expected %q but got %q", tc, output[i])
		}
	}
}

func TestItemsWithZeroQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"regular item", 5, 0},
		{"Sulfuras item", 5, 0},
		{"backstage passes item", 5, 0},
		{"Aged Brie item", 5, 0},
		{"Conjured item", 5, 0},
	}

	var expected = []*gildedrose.Item{
		{"regular item", 4, 0},
		{"Sulfuras item", 5, 0},
		{"backstage passes item", 4, 3},
		{"Aged Brie item", 4, 1},
		{"Conjured item", 4, 0},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {

		if strconv.Itoa(tc.Quality) != strconv.Itoa(items[i].Quality) {
			t.Errorf("Expected %q %q but got %q", items[i].Name, strconv.Itoa(tc.Quality), strconv.Itoa(items[i].Quality))
		}
		if strconv.Itoa(tc.SellIn) != strconv.Itoa(items[i].SellIn) {
			t.Errorf("Expected %q but got %q", strconv.Itoa(tc.SellIn), strconv.Itoa(items[i].SellIn))
		}
	}
}

func OutputInventory(items []*gildedrose.Item) []string {
	days := 2
	output := make([]string, 0)

	for day := 0; day < days; day++ {
		for _, item := range items {
			itemAsStr := item.Name + " " + strconv.Itoa(item.SellIn) + " " + strconv.Itoa(item.Quality)
			output = append(output, itemAsStr)
		}
		gildedrose.UpdateInventory(items)
	}
	return output
}
