package business

import (
	"dream/building"
	"testing"
)

func TestBusiness(t *testing.T) {
	convenienceStore := NewConvenienceStore(1)
	if convenienceStore.Type != building.ConvenienceStore && convenienceStore.Class != building.Business {
		t.Fail()
	}

	school := NewSchool(1)
	if school.Type != building.School && school.Class != building.Business {
		t.Fail()
	}

	couture := NewCouture(1)
	if couture.Type != building.Couture && couture.Class != building.Business {
		t.Fail()
	}

	hardwareStore := NewHardwareStore(1)
	if hardwareStore.Type != building.HardwareStore && hardwareStore.Class != building.Business {
		t.Fail()
	}

	foodMarket := NewFoodMarket(1)
	if foodMarket.Type != building.FoodMarket && foodMarket.Class != building.Business {
		t.Fail()
	}

	bookCity := NewBookCity(1)
	if bookCity.Type != building.BookCity && bookCity.Class != building.Business {
		t.Fail()
	}

	businessCentre := NewBusinessCentre(1)
	if businessCentre.Type != building.BusinessCentre && businessCentre.Class != building.Business {
		t.Fail()
	}

	gasStation := NewGasStation(1)
	if gasStation.Type != building.GasStation && gasStation.Class != building.Business {
		t.Fail()
	}

	folkFood := NewFolkFood(1)
	if folkFood.Type != building.FolkFood && folkFood.Class != building.Business {
		t.Fail()
	}

	mediaVoice := NewMediaVoice(1)
	if mediaVoice.Type != building.MediaVoice && mediaVoice.Class != building.Business {
		t.Fail()
	}

}
