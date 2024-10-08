package gildedrose_test

import (
	"strconv"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestItemsQualityShouldNotFallBelowZero(t *testing.T) {
	var items = []*gildedrose.Item{
		{"regular item", 5, 0},
		{"Conjured item", 5, 0},
	}

	var expected = []*gildedrose.Item{
		{"regular item", 4, 0},
		{"Conjured item", 4, 0},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func TestItemsQualityShouldNotExceedMax(t *testing.T) {
	var items = []*gildedrose.Item{
		{"regular item", 5, 50},
		{"backstage passes item", 5, 50},
		{"backstage passes item 2", 9, 48},
		{"Aged Brie item", 5, 50},
		{"Conjured item", 5, 50},
	}

	var expected = []*gildedrose.Item{
		{"regular item", 4, 49},
		{"backstage passes item", 4, 50},
		{"backstage passes item 2", 8, 50},
		{"Aged Brie item", 4, 50},
		{"Conjured item", 4, 48},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func TestSulfurasItemIsConstant(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras item", 5, 20},
	}

	var expected = []*gildedrose.Item{
		{"Sulfuras item", 5, 80},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func TestBackstageItemsUpdatesAsExpected(t *testing.T) {
	var items = []*gildedrose.Item{
		{"backstage passes item", 15, 22},
		{"backstage passes item 2", 9, 22},
		{"backstage passes item 3", 3, 22},
		{"backstage passes item 4", 0, 22},
	}

	var expected = []*gildedrose.Item{
		{"backstage passes item", 14, 23},
		{"backstage passes item 2", 8, 24},
		{"backstage passes item 3", 2, 25},
		{"backstage passes item 4", -1, 0},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func TestAgedBrieItemsUpdatesAsExpected(t *testing.T) {
	var items = []*gildedrose.Item{
		{"aged brie item", 5, 22},
		{"aged brie item", 0, 50},
		{"aged brie item", -3, 0},
	}

	var expected = []*gildedrose.Item{
		{"aged brie item", 4, 23},
		{"aged brie item", -1, 50},
		{"aged brie item", -4, 1},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func TestRegularItemsDecreaseDoublePastSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{"regular item", 0, 22},
		{"regular item 2", -2, 10},
	}

	var expected = []*gildedrose.Item{
		{"regular item", -1, 20},
		{"regular item 2", -3, 8},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func TestConjuredItemsDecreaseDouble(t *testing.T) {
	var items = []*gildedrose.Item{
		{"conjured item", -2, 50},
		{"conjured item 2", 0, 0},
		{"conjured item 3", 5, 23},
	}

	var expected = []*gildedrose.Item{
		{"conjured item", -3, 48},
		{"conjured item 2", -1, 0},
		{"conjured item 3", 4, 21},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func checkForEquality(t *testing.T, tc *gildedrose.Item, item *gildedrose.Item) {
	if strconv.Itoa(tc.Quality) != strconv.Itoa(item.Quality) {
		t.Errorf(" %q quality: Expected %q but got %q", item.Name, strconv.Itoa(tc.Quality), strconv.Itoa(item.Quality))
	}
	if strconv.Itoa(tc.SellIn) != strconv.Itoa(item.SellIn) {
		t.Errorf("SellIn: Expected %q but got %q", strconv.Itoa(tc.SellIn), strconv.Itoa(item.SellIn))
	}
}

func TestCharacterisedTestAutomated(t *testing.T) {
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

	output := characterisedTestAutomated(items)

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

func characterisedTestAutomated(items []*gildedrose.Item) []string {
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
