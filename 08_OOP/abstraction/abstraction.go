package abstraction

import "fmt"

// Vehicle interface demonstrating abstraction
type Vehicle interface {
	Start() string
	Stop() string
}

// Car struct implementing Vehicle interface
type Car struct {
	Model string
}

func (c Car) Start() string {
	return fmt.Sprintf("%s is starting", c.Model)
}

func (c Car) Stop() string {
	return fmt.Sprintf("%s is stopping", c.Model)
}

// Bike struct implementing Vehicle interface
type Bike struct {
	Model string
}

func (b Bike) Start() string {
	return fmt.Sprintf("%s is starting", b.Model)
}

func (b Bike) Stop() string {
	return fmt.Sprintf("%s is stopping", b.Model)
}

func main() {
	var v Vehicle

	v = Car{Model: "Toyota"}
	fmt.Println(v.Start())
	fmt.Println(v.Stop())

	v = Bike{Model: "Yamaha"}
	fmt.Println(v.Start())
	fmt.Println(v.Stop())
}
