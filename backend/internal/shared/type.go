package shared

type ProductQueryParams struct {
	Search string
	Filter string
	Page   int
	Limit  int
	Offset int
}

var ValidFilters = map[string]bool{
	"all":       true,
	"active":    true,
	"inactive":  true,
	"low_stock": true,
}
