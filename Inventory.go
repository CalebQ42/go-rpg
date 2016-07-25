package rpg

//Inventory is a collection of items ([]Item type)
type Inventory []Item

//HasItem returns if the Inventory has an item with the given name is present in
//the inventory
func (i Inventory) HasItem(name string) bool {
	for _, v := range i {
		if v.Name == name {
			return true
		}
	}
	return false
}

//AddItem adds an item to the inventory
func (i *Inventory) AddItem(item Item) {
	*i = append(*i, item)
}

//Items returns an string array with the names of the items in the inventory
func (i Inventory) Items() (out []string) {
	for _, v := range i {
		out = append(out, v.Name)
	}
	return
}

//RemoveItem removes the item with the given name from the Inventory. Removes only
//the first instance.
func (i Inventory) RemoveItem(name string) {
	for j, v := range i {
		if v.Name == name {
			i = append(i[:j], i[j+1:]...)
			return
		}
	}
}
