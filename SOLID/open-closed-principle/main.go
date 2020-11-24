package open_closed_principle

type Color int
type Size int

const (
	red Color = iota
	green
	blue
)

type Product struct {
	name string
	color Color
	size Size
}

type Filter struct {
	// ...some settings
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products{
		if v.color == color{
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
}
