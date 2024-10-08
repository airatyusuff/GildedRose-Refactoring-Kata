package gildedrose_test

import (
	"strconv"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_ItemsQualityShouldNotFallBelowZero(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Elixir of the Mongoose", 5, 0},
		{"Conjured Mana Cake", 5, 0},
	}

	var expected = []*gildedrose.Item{
		{"Elixir of the Mongoose", 4, 0},
		{"Conjured Mana Cake", 4, 0},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func Test_ItemsQualityShouldNotExceedMax(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 5, 50},
		{"Elixir of the Mongoose", 5, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 48},
		{"Conjured Mana Cake", 5, 50},
	}

	var expected = []*gildedrose.Item{
		{"Aged Brie", 4, 50},
		{"Elixir of the Mongoose", 4, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 50},
		{"Conjured Mana Cake", 4, 48},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func Test_SulfurasItemQualityIsConstant(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 5, 80},
	}

	var expected = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 5, 80},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func Test_BackstageItemsUpdatesAsExpected(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 22},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 22},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 22},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 22},
	}

	var expected = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 14, 23},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 24},
		{"Backstage passes to a TAFKAL80ETC concert", 2, 25},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func Test_AgedBrieItemsUpdatesAsExpected(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 5, 22},
		{"Aged Brie", 0, 50},
		{"Aged Brie", -3, 0},
	}

	var expected = []*gildedrose.Item{
		{"Aged Brie", 4, 23},
		{"Aged Brie", -1, 50},
		{"Aged Brie", -4, 1},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func Test_RegularItemsDecreaseDoublePastSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Elixir of the Mongoose", 0, 22},
		{"+5 Dexterity Vest", -2, 10},
	}

	var expected = []*gildedrose.Item{
		{"Elixir of the Mongoose", -1, 20},
		{"+5 Dexterity Vest", -3, 8},
	}

	gildedrose.UpdateInventory(items)

	for i, tc := range expected {
		checkForEquality(t, tc, items[i])
	}
}

func Test_ConjuredItemsDecreaseDouble(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", -2, 50},
		{"Conjured Mana Cake", 0, 0},
		{"Conjured Mana Cake", 5, 23},
	}

	var expected = []*gildedrose.Item{
		{"Conjured Mana Cake", -3, 48},
		{"Conjured Mana Cake", -1, 0},
		{"Conjured Mana Cake", 4, 21},
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
		"+5 Dexterity Vest 10 20",
		"Aged Brie 2 0",
		"Elixir of the Mongoose 5 0",
		"Sulfuras, Hand of Ragnaros 0 80",
		"Sulfuras, Hand of Ragnaros -1 80",
		"Backstage passes to a TAFKAL80ETC concert 15 20",
		"Backstage passes to a TAFKAL80ETC concert 10 50",
		"Backstage passes to a TAFKAL80ETC concert 0 20",
		"Conjured Mana Cake 3 6",
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
