package rpg

//Item is a basic item in an inventory
type Item struct {
	Name        string
	Description string
}

//CreateItem creates an item
func CreateItem(name, description string) (out Item) {
	out.Name = name
	out.Description = description
	return
}
