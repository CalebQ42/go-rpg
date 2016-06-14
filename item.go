package rpg

//Item TODO
type Item struct {
	name, desc string
}

//CreateItem TODO
func CreateItem(name, description string) (out Item) {
	out.name = name
	out.desc = description
	return
}

//Name TODO
func (i Item) Name() string {
	return i.name
}

//Description TODO
func (i Item) Description() string {
	return i.desc
}
