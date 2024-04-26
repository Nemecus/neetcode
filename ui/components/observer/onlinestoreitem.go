package observer

type item struct {
	customerList []observer
	name         string
	stock        int
}

func newItem(name string, stock int) *item {
	return &item{
		name:  name,
		stock: stock,
	}
}

func (i *item) updateStock(stockCount int) {
	oldStock := i.stock
	i.stock = stockCount

	if oldStock == 0 && stockCount > 0 {
		i.notifyAll()
	}
}

func (i *item) subscribe(o observer) {
	i.customerList = append(i.customerList, o)
}

func (i *item) unsubscribe(o observer) {
	i.customerList = removeFromList(i.customerList, o)
}

func (i *item) notifyAll() {
	for _, customer := range i.customerList {
		customer.notify(i.name)
	}
}

func removeFromList(customerList []observer, customerToRemove observer) []observer {
	customerCount := len(customerList)
	for i, customer := range customerList {
		if customerToRemove.getName() == customer.getName() {
			customerList[customerCount-1], customerList[i] = customerList[i], customerList[customerCount-1]
			return customerList[:customerCount-1]
		}
	}
	return customerList
}
