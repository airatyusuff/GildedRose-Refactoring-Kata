package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestOutputDaysInventory(t *testing.T) {
	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 20},
		{"Conjured Mana Cake", 3, 6},
	}

	output := gildedrose.OutputDaysInventory(items)

	expected := []string{
		"+5 Dexterity Vest 10 20",
		"Aged Brie 2 0",
		"Elixir of the Mongoose 5 7",
		"Sulfuras, Hand of Ragnaros 0 80",
		"Sulfuras, Hand of Ragnaros -1 80",
		"Backstage passes to a TAFKAL80ETC concert 15 20",
		"Backstage passes to a TAFKAL80ETC concert 10 50",
		"Backstage passes to a TAFKAL80ETC concert 4 20",
		"Conjured Mana Cake 3 6",
		"+5 Dexterity Vest 9 19",
		"Aged Brie 1 1",
		"Elixir of the Mongoose 4 6",
		"Sulfuras, Hand of Ragnaros 0 80",
		"Sulfuras, Hand of Ragnaros -1 80",
		"Backstage passes to a TAFKAL80ETC concert 14 21",
		"Backstage passes to a TAFKAL80ETC concert 9 50",
		"Backstage passes to a TAFKAL80ETC concert 3 23",
		"Conjured Mana Cake 2 5",
	}

	for i, tc := range expected {
		if tc != output[i] {
			t.Errorf("Expected %q but got %q", tc, output[i])
		}
	}
}

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "foo" {
		t.Errorf("Name: Expected %s but got %s ", "foo", items[0].Name)
	}
}
