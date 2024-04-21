package factorymethod

//car
//truck
//bike

type Vehicle interface {
	GetType() string
}

type Car struct{}
type Truck struct{}
type Bike struct{}

func (c *Car) GetType() string {
	return "Car"
}

func (t *Truck) GetType() string {
	return "Truck"
}

func (b *Bike) GetType() string {
	return "Bike"
}

type VehicleFactory interface {
	CreateVehicle() Car
}

type CarFactory struct{}
type TruckFactory struct{}
type BikeFactory struct{}

func (cf *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}

func (tf *TruckFactory) CreateVehicle() Vehicle {
	return &Truck{}
}

func (bf *BikeFactory) CreateVehicle() Vehicle {
	return &Bike{}
}
