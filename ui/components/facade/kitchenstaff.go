package facade

type KitchenStaff struct {
}

func (ks *KitchenStaff) packageOrder(packagedFood *PackagedFood, food *Food) *Food {
	return packagedFood.newPackagedFood(food)
}
