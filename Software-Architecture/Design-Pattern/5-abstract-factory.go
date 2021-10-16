package designpattern

import (
	"fmt"
	"os"
)

var (
	MODERN    = "modern"
	VICTORIAN = "victorian"
	ART       = "art"
)

// Perusahaan Logistik
// Product ---> Kapal, Truck, Psawat
// Factory ----> menghasilkan Kapal, Truck, Pesawat

// service ---> sameDay, OneDay, Regular

// FactorySameDay ---> (Kapal, Truck, Pesawat) sameday
// FactoryOneDay ---> (Kapal, Truck, Pesawat) OneDay
// FactoryRegular ---> (Kapal, Truck, Pesawat) Regular

// interface untuk factory per-product ==================
type Chair interface {
	sitOn()
}

type Sofa interface {
	PeopleSitOn()
}

type CoffeTable interface {
	hasCoffeeOn()
}

// define types for each product ==================

// chair types ==================
type ModernChair struct{}

func (mc *ModernChair) sitOn() {
	fmt.Println("Someone siton modern chair")
}

type VictorianChair struct{}

func (vc *VictorianChair) sitOn() {
	fmt.Println("Someone siton victorian chair")
}

type ArtChair struct{}

func (ac *ArtChair) sitOn() {
	fmt.Println("Someone siton art chair")
}

// Sofa types ==================
type ModernSofa struct{}

func (ms *ModernSofa) PeopleSitOn() {
	fmt.Println("Some People Sit On modern Sofa")
}

type VictorianSofa struct{}

func (vs *VictorianSofa) PeopleSitOn() {
	fmt.Println("Some People Sit On victorian Sofa")
}

type ArtSofa struct{}

func (as *ArtSofa) PeopleSitOn() {
	fmt.Println("Some People Sit On art Sofa")
}

// CoffeTable types ==================
type ModernCoffeTable struct{}

func (mct *ModernCoffeTable) hasCoffeeOn() {
	fmt.Println("it has an Coffee On modern CoffeTable")
}

type VictorianCoffeTable struct{}

func (vct *VictorianCoffeTable) hasCoffeeOn() {
	fmt.Println("it has an Coffee On victorian CoffeTable")
}

type ArtCoffeTable struct{}

func (act *ArtCoffeTable) hasCoffeeOn() {
	fmt.Println("it has an Coffee On art CoffeTable")
}

// interface untuk factory-factory nya ==================
type FurnitureFactory interface {
	CreteChair() Chair
	CreteSofa() Sofa
	CreteCoffeTable() CoffeTable
}

// factory-factory ==================

// modern-factory ==================
type ModernFurnitureFactory struct{}

func (mff *ModernFurnitureFactory) CreteChair() Chair {
	return &ModernChair{}
}

func (mff *ModernFurnitureFactory) CreteSofa() Sofa {
	return &ModernSofa{}
}

func (mff *ModernFurnitureFactory) CreteCoffeTable() CoffeTable {
	return &ModernCoffeTable{}
}

// Victorian-factory ==================
type VictorianFurnitureFactory struct{}

func (mff *VictorianFurnitureFactory) CreteChair() Chair {
	return &VictorianChair{}
}

func (mff *VictorianFurnitureFactory) CreteSofa() Sofa {
	return &VictorianSofa{}
}

func (mff *VictorianFurnitureFactory) CreteCoffeTable() CoffeTable {
	return &VictorianCoffeTable{}
}

// Art-factory ==================
type ArtFurnitureFactory struct{}

func (mff *ArtFurnitureFactory) CreteChair() Chair {
	return &ArtChair{}
}

func (mff *ArtFurnitureFactory) CreteSofa() Sofa {
	return &ArtSofa{}
}

func (mff *ArtFurnitureFactory) CreteCoffeTable() CoffeTable {
	return &ArtCoffeTable{}
}

func AbstractFactory() {
	modernFactory := &ModernFurnitureFactory{}
	victorianFactory := &VictorianFurnitureFactory{}
	artFactory := &ArtFurnitureFactory{}

	var factory FurnitureFactory

	varian := os.Args[1]

	switch varian {
	case MODERN:
		factory = modernFactory
	case VICTORIAN:
		factory = victorianFactory
	case ART:
		factory = artFactory
	}

	chair := factory.CreteChair()
	chair.sitOn()

	sofa := factory.CreteSofa()
	sofa.PeopleSitOn()

	coffeTable := factory.CreteCoffeTable()
	coffeTable.hasCoffeeOn()
}
