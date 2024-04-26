package observer

type observer interface {
	notify(string)
	getName() string
	countNotifications() int
}
