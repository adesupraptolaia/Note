package designpattern

import "fmt"

type Logistik interface {
	// qty int
	deliver()
}

// =========== TRUCK ==========
type Truck struct {
	qty int
}

func (t *Truck) deliver() {
	fmt.Printf("kirim %d barang menggunakan truk\n", t.qty)
}

// =========== Ship ==========
type Ship struct {
	qty int
}

func (t *Ship) deliver() {
	fmt.Printf("kirim %d barang menggunakan kapal\n", t.qty)
}

// CREATOR
type factory struct{}

func (f *factory) createLogistik(vehicleType string, qty int) Logistik {
	if vehicleType == "truck" {
		return &Truck{qty}
	} else if vehicleType == "ship" {
		return &Ship{qty}
	}
	return nil
}

func FactoryMethod() {
	f := factory{}

	truck1 := f.createLogistik("truck", 250)
	truck1.deliver()

	ship1 := f.createLogistik("ship", 250)
	ship1.deliver()
}
