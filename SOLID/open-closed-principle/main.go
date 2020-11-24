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

func main() {

}
