package rpg

//Item TODO
type Item struct {
	Name        string
	Description string
}

//CreateItem TODO
func CreateItem(name, description string) (out Item) {
	out.Name = name
	out.Description = description
	return
}
