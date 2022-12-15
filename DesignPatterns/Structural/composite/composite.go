package composite

type Composite interface {
	SearchName(phrase string) (bool, string)
}
