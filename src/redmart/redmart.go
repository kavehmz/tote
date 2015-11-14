package redmart

import (
	"fmt"
	"sort"
)

// Product defines Redmart products
type Product struct {
	ID     int
	Price  int
	Length int
	Width  int
	Height int
	Weight int
}

func (p Product) cubicValue() float64 {
	return float64(p.Price) / float64(p.volume())
}
func (p Product) volume() int {
	return p.Length * p.Width * p.Height
}

// For sorting based on cubic value
type products []Product

func (ps products) Len() int {
	return len(ps)
}
func (ps products) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func (ps products) Less(i, j int) bool {
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

func findMax(n int, maxVolume int, ps []Product) int {
	var val int
	if n == 0 || maxVolume <= 0 {
		return 0
	}
	if ps[n-1].volume() > maxVolume {
		val = findMax(n-1, maxVolume, ps)
	} else {
		a := findMax(n-1, maxVolume, ps)
		b := findMax(n-1, maxVolume-ps[n-1].volume(), ps) + ps[n-1].Price
		if a > b {
			val = a
		} else {
			val = b
		}
	}
	return val
}

func printElements(s [][]bool, ps []Product, v int, n int) {
	if v == 0 || n == 0 {
		return
	}
	if s[n][v] {
		fmt.Println(ps[n-1])
		printElements(s, ps, v-ps[n-1].volume(), n-1)
	} else {
		printElements(s, ps, v, n-1)
	}
}

// HighestValue will calculate highest value of products that is choosable from Redmart products
func HighestValue(tote Tote, ps []Product) string {
	ps = removeUnfittables(tote, ps)
	sort.Sort(products(ps))

	fmt.Println(len(ps))
	n := len(ps)

	w := make([]int, n+1)
	v := make([]int, n+1)
	for k, val := range ps {
		w[k+1] = val.volume()
		v[k+1] = val.Price
	}

	m := make([][]int, n+1)
	s := make([][]bool, n+1)
	for k := range m {
		m[k] = make([]int, tote.volume()+1)
		s[k] = make([]bool, tote.volume()+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= tote.volume(); j++ {
			if w[i] <= j {
				if m[i-1][j] >= m[i-1][j-w[i]]+v[i] {
					m[i][j] = m[i-1][j]
					s[i][j] = false
				} else {
					m[i][j] = m[i-1][j-w[i]] + v[i]
					s[i][j] = true
				}
			} else {
				m[i][j] = m[i-1][j]
				s[i][j] = false
			}

		}
	}

	fmt.Println(m[n][tote.volume()])

	printElements(s, ps, tote.volume(), n)

	return "t"
}
