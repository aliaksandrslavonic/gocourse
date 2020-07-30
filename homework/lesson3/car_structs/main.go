package main

import "fmt"

func main() {
	car1 := vehicle{Brand: "Honda", Year: 2015, TrunkVolume: 100, EngineOn: false, WindowsOpen: false, TrunkFullness: 30}
	car3 := vehicle{Brand: "Suzuki", Year: 2013, TrunkVolume: 80, EngineOn: true, WindowsOpen: false, TrunkFullness: 10}
	car2 := truck{General: vehicle{Brand: "Volvo", Year: 2010, TrunkVolume: 20000, EngineOn: false, WindowsOpen: true, TrunkFullness: 150}, Mass: 5}

	fmt.Println("Print structs")
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car3)

	fmt.Println("Print brands")
	fmt.Println(car1.Brand)
	fmt.Println(car2.General.Brand)

	fmt.Println("car1 = car3")
	fmt.Println(car1 == car3)

	fmt.Println("Check Engine")
	fmt.Println(car3.EngineOn)
	fmt.Println("EngineOFF")
	car3.EngineOn = false
	fmt.Println(car3.EngineOn)

	fmt.Println("Change Brand")
	fmt.Println("before ", car1.Brand)
	car1 = changeBrand(car1, "Fiat")
	fmt.Println("after", car1.Brand)

}

type vehicle struct {
	Brand         string
	Year          int
	TrunkVolume   int
	EngineOn      bool
	WindowsOpen   bool
	TrunkFullness int
}

type truck struct {
	General vehicle
	Mass    int
}

func changeBrand(car vehicle, newBrand string) vehicle {
	car.Brand = newBrand
	return car
}
