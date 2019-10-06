package industry

import (
	"dream/building"
	"testing"
)

func TestIndustry(t *testing.T) {
	woodFactory := NewWoodFactory(1)
	if woodFactory.Type != building.WoodFactory && woodFactory.Class != building.Industry {
		t.Fail()
	}
	paperMill := NewPaperMill(1)
	if paperMill.Type != building.PaperMill && paperMill.Class != building.Industry {
		t.Fail()
	}
	waterworks := NewWaterworks(1)
	if waterworks.Type != building.Waterworks && waterworks.Class != building.Industry {
		t.Fail()
	}
	powerPlant := NewPowerPlant(1)
	if powerPlant.Type != building.PowerPlant && powerPlant.Class != building.Industry {
		t.Fail()
	}
	foodManufacturer := NewFoodManufacturer(1)
	if foodManufacturer.Type != building.FoodManufacturer && foodManufacturer.Class != building.Industry {
		t.Fail()
	}
	ironworks := NewIronworks(1)
	if ironworks.Type != building.Ironworks && ironworks.Class != building.Industry {
		t.Fail()
	}
	textileMill := NewTextileMill(1)
	if textileMill.Type != building.TextileMill && textileMill.Class != building.Industry {
		t.Fail()
	}
	sparePartsFactory := NewSparePartsFactory(1)
	if sparePartsFactory.Type != building.SparePartsFactory && sparePartsFactory.Class != building.Industry {
		t.Fail()
	}
	penguinMachinery := NewPenguinMachinery(1)
	if penguinMachinery.Type != building.PenguinMachinery && penguinMachinery.Class != building.Industry {
		t.Fail()
	}
	petroleum := NewPetroleum(1)
	if petroleum.Type != building.Petroleum && petroleum.Class != building.Industry {
		t.Fail()
	}

}
