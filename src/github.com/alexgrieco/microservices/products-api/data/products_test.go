package data

import (
	"testing"
)

func TestChecksValidation(t*testing.T) {
	p := &Product{
		Name: "Alex",
		Price: 1.00,
		SKU: "abs-abc-dex",
	}
	
	err := p.Validate()
	
	if err != nil {
		t.Fatal(err)
	}
}