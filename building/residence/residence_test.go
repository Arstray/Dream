package residence

import (
	"dream/building"
	"testing"
)

func TestResidence(t *testing.T) {

	woodenHouse := NewWoodenHouse(1)
	if woodenHouse.Type != building.WoodenHouse && woodenHouse.Class != building.Residence {
		t.Fail()
	}
	steelStructure := NewSteelStructure(1)
	if steelStructure.Type != building.SteelStructure && steelStructure.Class != building.Residence {
		t.Fail()
	}
	bungalow := NewBungalow(1)
	if bungalow.Type != building.Bungalow && bungalow.Class != building.Residence {
		t.Fail()
	}
	studioApartment := NewStudioApartment(1)
	if studioApartment.Type != building.StudioApartment && studioApartment.Class != building.Residence {
		t.Fail()
	}
	residentialBuilding := NewResidentialBuilding(1)
	if residentialBuilding.Type != building.ResidentialBuilding && residentialBuilding.Class != building.Residence {
		t.Fail()
	}
	talentApartment := NewTalentApartment(1)
	if talentApartment.Type != building.TalentApartment && talentApartment.Class != building.Residence {
		t.Fail()
	}
	gardenHouse := NewGardenHouse(1)
	if gardenHouse.Type != building.GardenHouse && gardenHouse.Class != building.Residence {
		t.Fail()
	}
	chineseSmallHouse := NewChineseSmallHouse(1)
	if chineseSmallHouse.Type != building.ChineseSmallHouse && chineseSmallHouse.Class != building.Residence {
		t.Fail()
	}
	villa := NewVilla(1)
	if villa.Type != building.Villa && villa.Class != building.Residence {
		t.Fail()
	}
	renaissanceMansion := NewRenaissanceMansion(1)
	if renaissanceMansion.Type != building.RenaissanceMansion && renaissanceMansion.Class != building.Residence {
		t.Fail()
	}
}
