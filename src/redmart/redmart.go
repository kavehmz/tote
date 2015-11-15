// Package prime provides functionality to pack a tote with most value optimized
// products.
// This is basically a classic knapsack algorithm with addition of preferring
// lighter tote if there is no value loss.
//
// Value here is price of products to pick
package redmart

import (
	"fmt"
	"sort"
)

// Product defines a Redmart products
type Product struct {
	ID     int
	Price  int
	Length int
	Width  int
	Height int
	Weight int
}

func (p Product) volume() int {
	return p.Length * p.Width * p.Height
}

// For sorting based on cubic value
type productsByWeight []Product

func (ps productsByWeight) Len() int {
	return len(ps)
}
func (ps productsByWeight) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func (ps productsByWeight) Less(i, j int) bool {
	return ps[i].Weight < ps[j].Weight
}

// Tote will define dimentions of a Tote
type Tote struct {
	Length int
	Width  int
	Height int
}

func (t Tote) volume() int {
	return t.Length * t.Width * t.Height
}

func removeUnfittables(tote Tote, ps []Product) []Product {
	for k, v := range ps {
		if v.Length > tote.Length || v.Width > tote.Width || v.Height > tote.Height {
			ps = append(ps[:k], ps[k+1:]...)
		}
	}
	return ps
}

var selection []Product

func findSelection(s [][]bool, ps []Product, v int, n int) {
	if v == 0 || n == 0 {
		return
	}
	if s[n][v] {
		selection = selection[:len(selection)+1]
		selection[len(selection)-1] = ps[n-1]
		findSelection(s, ps, v-ps[n-1].volume(), n-1)
	} else {
		findSelection(s, ps, v, n-1)
	}
}

func decisionMaps(tote Tote, ps []Product) ([][]int, [][]bool) {
	n := len(ps)
	volumes := make([]int, n+1)
	prices := make([]int, n+1)
	for k, val := range ps {
		volumes[k+1] = val.volume()
		prices[k+1] = val.Price
	}

	valMap := make([][]int, n+1)
	selMap := make([][]bool, n+1)
	for k := range valMap {
		valMap[k] = make([]int, tote.volume()+1)
		selMap[k] = make([]bool, tote.volume()+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= tote.volume(); j++ {
			if volumes[i] <= j {
				if valMap[i-1][j] >= valMap[i-1][j-volumes[i]]+prices[i] {
					valMap[i][j] = valMap[i-1][j]
					selMap[i][j] = false
				} else {
					valMap[i][j] = valMap[i-1][j-volumes[i]] + prices[i]
					selMap[i][j] = true
				}
			} else {
				valMap[i][j] = valMap[i-1][j]
				selMap[i][j] = false
			}

		}
	}
	return valMap, selMap
}

// PackTote will find the best combination of products to pick
func PackTote(tote Tote, ps []Product) {
	ps = removeUnfittables(tote, ps)
	sort.Sort(productsByWeight(ps))

	n := len(ps)
	selection = make([]Product, 0, n)
	valMap, selMap := decisionMaps(tote, ps)
	findSelection(selMap, ps, tote.volume(), n)

	var sum int
	for _, v := range selection {
		sum += v.ID
	}
	fmt.Println("Optimized value: ", valMap[n][tote.volume()])
	fmt.Println("Sum of product IDs: ", sum)
}
