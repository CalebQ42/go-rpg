package rpg

//Item TODO
type Item struct {
	Name, Desc string
}

//CreateItem TODO
func CreateItem(name, description string) (out Item) {
	out.Name = name
	out.Desc = description
	return
}

//Description TODO
func (i Item) Description() string {
	return i.Desc
}
