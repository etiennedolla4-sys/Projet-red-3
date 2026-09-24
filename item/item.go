package item

type Item interface {
	Name() string
	Description() string
	Type() string
}
