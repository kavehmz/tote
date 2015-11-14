package redmart

import "fmt"

// Product defines Redmart products
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

// Tote will define dimentions of a Tote
type Tote struct {
	Length int
	Width  int
	Height int
}

func (t Tote) volume() int {
	return t.Length * t.Width * t.Height
}

// HighestValue will calculate highest value of products that is choosable from Redmart products
func HighestValue(tote Tote, ps []Product) string {
	fmt.Println(len(ps))
	n := len(ps)

	w := make([]int, n+1)
	v := make([]int, n+1)
	for k, val := range ps {
		w[k+1] = val.volume()
		v[k+1] = val.Price
	}

	m := make([][]int, n+1)
	for k, _ := range m {
		m[k] = make([]int, tote.volume()+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= tote.volume(); j++ {
			if w[i] <= j {
				if m[i-1][j] >= m[i-1][j-w[i]]+v[i] {
					m[i][j] = m[i-1][j]
				} else {
					m[i][j] = m[i-1][j-w[i]] + v[i]
				}
			} else {
				m[i][j] = m[i-1][j]
			}

		}
	}

	fmt.Println(m[n][tote.volume()])
	return "t"
}
