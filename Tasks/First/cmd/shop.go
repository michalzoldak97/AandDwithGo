package main

import "strings"

type product struct {
	name   string
	amount int
	price  float64
}

type shop struct {
	products map[string][]product
}

func (s *shop) getSections() []string {
	sections := make([]string, len(s.products))

	i := 0
	for s := range s.products {
		sections[i] = s
		i++
	}

	return sections
}

func (s *shop) getProduct(pname string) (bool, product) {
	searchPName := strings.ToLower(pname)

	for section := range s.products {
		for _, prod := range s.products[section] {
			if strings.ToLower(prod.name) == searchPName {
				return true, prod
			}
		}
	}

	return false, product{}
}

func (s *shop) buyProduct(pname string, amount int) bool {
	for section := range s.products {
		for i, prod := range s.products[section] {
			if prod.name == pname {
				s.products[section][i].amount -= amount
				return true
			}
		}
	}

	return false
}

func newShop() shop {
	return shop{
		products: map[string][]product{"bread": {{
			name:   "Bread",
			amount: 10,
			price:  5.45,
		}, {
			name:   "Roll",
			amount: 24,
			price:  1.19,
		}}, "alko": {{
			name:   "Harnas",
			amount: 999,
			price:  2.,
		}, {
			name:   "Apple Blossom",
			amount: 17,
			price:  9.59,
		}, {
			name:   "Johny Walker",
			amount: 6,
			price:  119.88,
		}}},
	}
}
