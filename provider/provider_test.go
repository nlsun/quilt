package provider

import (
	"testing"

	"github.com/NetSys/quilt/dsl"
)

func TestConstraints(t *testing.T) {
	checkConstraint := func(descriptions []description, ram dsl.Range, cpu dsl.Range,
		maxPrice float64, exp string) {
		resSize := pickBestSize(descriptions, ram, cpu, maxPrice)
		if resSize != exp {
			t.Errorf("bad size picked. Expected %s, got %s", exp, resSize)
		}
	}

	// Test all constraints specified with valid price
	testDescriptions := []description{
		{size: "size1", price: 2, ram: 2, cpu: 2},
	}
	checkConstraint(testDescriptions, dsl.Range{Min: 1, Max: 3},
		dsl.Range{Min: 1, Max: 3}, 2, "size1")

	// Test no max
	checkConstraint(testDescriptions, dsl.Range{Min: 1},
		dsl.Range{Min: 1}, 2, "size1")

	// Test exact match
	checkConstraint(testDescriptions, dsl.Range{Min: 2},
		dsl.Range{Min: 2}, 2, "size1")

	// Test no match
	checkConstraint(testDescriptions, dsl.Range{Min: 3},
		dsl.Range{Min: 2}, 2, "")

	// Test price too expensive
	checkConstraint(testDescriptions, dsl.Range{Min: 2},
		dsl.Range{Min: 2}, 1, "")

	// Test multiple matches (should pick cheapest)
	testDescriptions = []description{
		{size: "size2", price: 2, ram: 8, cpu: 4},
		{size: "size3", price: 1, ram: 4, cpu: 4},
		{size: "size4", price: 0.5, ram: 3, cpu: 4},
	}
	checkConstraint(testDescriptions, dsl.Range{Min: 4},
		dsl.Range{Min: 3}, 2, "size3")

	// Test infinite price
	checkConstraint(testDescriptions, dsl.Range{Min: 4},
		dsl.Range{Min: 3}, 0, "size3")

	// Test default ranges (should pick cheapest)
	checkConstraint(testDescriptions, dsl.Range{},
		dsl.Range{}, 0, "size4")

	// Test one default range (should pick only on the specified range)
	checkConstraint(testDescriptions, dsl.Range{Min: 4},
		dsl.Range{}, 0, "size3")
	checkConstraint(testDescriptions, dsl.Range{Min: 3},
		dsl.Range{}, 0, "size4")
}
