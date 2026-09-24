package item

func AddItem(inventory *[]Item, newItem Item, maxInventory int) bool {
	if len(*inventory) >= maxInventory {
		return false
	}

	*inventory = append(*inventory, newItem)
	return true
}

func RemoveItem(inventory *[]Item, index int) bool {
	if index < 0 || index >= len(*inventory) {
		return false
	}

	*inventory = append((*inventory)[:index], (*inventory)[index+1:]...)
	return true
}
