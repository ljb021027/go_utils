package format

type Options struct {
	Query   string
	ShowAll bool
	Page    int
	Array   []string
	ArrayA  []A
}

type A struct {
	Name string
	Age  int
}
