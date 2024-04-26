package observer

type iCustomer interface {
	subscribe(observer observer)
	unsubscribe(observer observer)
	notifyAll()
}
