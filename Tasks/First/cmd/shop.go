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
		}, {
			name:   "Cake",
			amount: 8,
			price:  18.99,
		}, {
			name:   "Doughnut",
			amount: 7,
			price:  3.11,
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
		}}, "dairy": {{
			name:   "Butter",
			amount: 31,
			price:  500.99,
		}, {
			name:   "Milk",
			amount: 72,
			price:  4.17,
		}, {
			name:   "Monte",
			amount: 12,
			price:  13.90,
		}, {
			name:   "Kefir",
			amount: 6,
			price:  7.88,
		},
		}, "tools": {{
			name:   "Glebogryzarka",
			amount: 1,
			price:  1785.33,
		}, {
			name:   "Klucz do glebogryzarki",
			amount: 1,
			price:  .99,
		},
		}},
	}
}
