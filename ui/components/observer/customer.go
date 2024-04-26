package observer

type customer struct {
	name          string
	notifications int
}

func (c *customer) notify(item string) {
	c.notifications++
}

func (c *customer) getName() string {
	return c.name
}

func (c *customer) countNotifications() int {
	return c.notifications
}
