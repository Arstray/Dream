package building

import (
	"testing"
)

func TestBusiness(t *testing.T) {
	convenienceStore := NewConvenienceStore(1)
	if convenienceStore.Type != ConvenienceStore && convenienceStore.Class != Business {
		t.Fail()
	}

	school := NewSchool(1)
	if school.Type != School && school.Class != Business {
		t.Fail()
	}

	couture := NewCouture(1)
	if couture.Type != Couture && couture.Class != Business {
		t.Fail()
	}

	hardwareStore := NewHardwareStore(1)
	if hardwareStore.Type != HardwareStore && hardwareStore.Class != Business {
		t.Fail()
	}

	foodMarket := NewFoodMarket(1)
	if foodMarket.Type != FoodMarket && foodMarket.Class != Business {
		t.Fail()
	}

	bookCity := NewBookCity(1)
	if bookCity.Type != BookCity && bookCity.Class != Business {
		t.Fail()
	}

	businessCentre := NewBusinessCentre(1)
	if businessCentre.Type != BusinessCentre && businessCentre.Class != Business {
		t.Fail()
	}

	gasStation := NewGasStation(1)
	if gasStation.Type != GasStation && gasStation.Class != Business {
		t.Fail()
	}

	folkFood := NewFolkFood(1)
	if folkFood.Type != FolkFood && folkFood.Class != Business {
		t.Fail()
	}

	mediaVoice := NewMediaVoice(1)
	if mediaVoice.Type != MediaVoice && mediaVoice.Class != Business {
		t.Fail()
	}

}

func TestResidence(t *testing.T) {

	woodenHouse := NewWoodenHouse(1)
	if woodenHouse.Type != WoodenHouse && woodenHouse.Class != Residence {
		t.Fail()
	}
	steelStructure := NewSteelStructure(1)
	if steelStructure.Type != SteelStructure && steelStructure.Class != Residence {
		t.Fail()
	}
	bungalow := NewBungalow(1)
	if bungalow.Type != Bungalow && bungalow.Class != Residence {
		t.Fail()
	}
	studioApartment := NewStudioApartment(1)
	if studioApartment.Type != StudioApartment && studioApartment.Class != Residence {
		t.Fail()
	}
	residentialBuilding := NewResidentialBuilding(1)
	if residentialBuilding.Type != ResidentialBuilding && residentialBuilding.Class != Residence {
		t.Fail()
	}
	talentApartment := NewTalentApartment(1)
	if talentApartment.Type != TalentApartment && talentApartment.Class != Residence {
		t.Fail()
	}
	gardenHouse := NewGardenHouse(1)
	if gardenHouse.Type != GardenHouse && gardenHouse.Class != Residence {
		t.Fail()
	}
	chineseSmallHouse := NewChineseSmallHouse(1)
	if chineseSmallHouse.Type != ChineseSmallHouse && chineseSmallHouse.Class != Residence {
		t.Fail()
	}
	villa := NewVilla(1)
	if villa.Type != Villa && villa.Class != Residence {
		t.Fail()
	}
	renaissanceMansion := NewRenaissanceMansion(1)
	if renaissanceMansion.Type != RenaissanceMansion && renaissanceMansion.Class != Residence {
		t.Fail()
	}
}
func TestIndustry(t *testing.T) {
	woodFactory := NewWoodFactory(1)
	if woodFactory.Type != WoodFactory && woodFactory.Class != Industry {
		t.Fail()
	}
	paperMill := NewPaperMill(1)
	if paperMill.Type != PaperMill && paperMill.Class != Industry {
		t.Fail()
	}
	waterworks := NewWaterworks(1)
	if waterworks.Type != Waterworks && waterworks.Class != Industry {
		t.Fail()
	}
	powerPlant := NewPowerPlant(1)
	if powerPlant.Type != PowerPlant && powerPlant.Class != Industry {
		t.Fail()
	}
	foodManufacturer := NewFoodManufacturer(1)
	if foodManufacturer.Type != FoodManufacturer && foodManufacturer.Class != Industry {
		t.Fail()
	}
	ironworks := NewIronworks(1)
	if ironworks.Type != Ironworks && ironworks.Class != Industry {
		t.Fail()
	}
	textileMill := NewTextileMill(1)
	if textileMill.Type != TextileMill && textileMill.Class != Industry {
		t.Fail()
	}
	sparePartsFactory := NewSparePartsFactory(1)
	if sparePartsFactory.Type != SparePartsFactory && sparePartsFactory.Class != Industry {
		t.Fail()
	}
	penguinMachinery := NewPenguinMachinery(1)
	if penguinMachinery.Type != PenguinMachinery && penguinMachinery.Class != Industry {
		t.Fail()
	}
	petroleum := NewPetroleum(1)
	if petroleum.Type != Petroleum && petroleum.Class != Industry {
		t.Fail()
	}

}
