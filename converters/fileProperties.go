package converters

type TagComponents struct {
	Name  string
	Color string
}

type FileProperties struct {
	Name string
	Tags []TagComponents
}
