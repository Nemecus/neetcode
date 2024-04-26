package facade

type driveThruFacade struct {
	kitchenStaff *KitchenStaff
	chef         *Chef
	cashier      *Cashier
	order        *Order
	food         *Food
	packagedFood *PackagedFood
}

func newDriveThruFacade() *driveThruFacade {
	dtFacade := &driveThruFacade{
		kitchenStaff: &KitchenStaff{},
		chef:         &Chef{},
		cashier:      &Cashier{},
		order:        &Order{},
		food:         &Food{},
		packagedFood: &PackagedFood{},
	}

	return dtFacade
}

func (dtf *driveThruFacade) takeOrder(orderContents string, takeOut bool) *Food {
	dtf.order.contents = orderContents
	dtf.order.takeOut = takeOut

	dtf.food = dtf.chef.prepareFood(dtf.order)

	if dtf.order.takeOut {
		return dtf.kitchenStaff.packageOrder(dtf.packagedFood, dtf.food)
	}
	return dtf.food
}
