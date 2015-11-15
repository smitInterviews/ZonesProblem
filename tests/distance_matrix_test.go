package testing

import (
	"fmt"
	"github.com/ZonesProblem/server"
	//"os"
	"testing"
)

func CreateZoneDataMap() map[string][]main.ZipCodeZoneData {
	var zoneDataMap = make(map[string][]main.ZipCodeZoneData)
	zoneDataMap["0"] = []main.ZipCodeZoneData{
		main.ZipCodeZoneData{
			0,
			1,
			1,
		},
		main.ZipCodeZoneData{
			2,
			3,
			2,
		},
		main.ZipCodeZoneData{
			5,
			6,
			3,
		},
		main.ZipCodeZoneData{
			7,
			9,
			4,
		},
		main.ZipCodeZoneData{
			12,
			15,
			5,
		},
	}
	return zoneDataMap
}

func TestExistingStartRangeElementSearch(t *testing.T) {
	zoneDataMap := CreateZoneDataMap()
	locator := &main.USPSZoneLocator{
		ZoneDataMap: zoneDataMap,
	}
	zone, err := locator.FindZone("0", 5)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if zone != 3 {
		fmt.Println("Incorrect zone returned. Zone: ", zone, " Expected: 3")
		t.Fail()
	}
}

func TestExistingEndRangeElementSearch(t *testing.T) {
	zoneDataMap := CreateZoneDataMap()
	locator := &main.USPSZoneLocator{
		ZoneDataMap: zoneDataMap,
	}
	zone, err := locator.FindZone("0", 9)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if zone != 4 {
		fmt.Println("Incorrect zone returned. Zone: ", zone, " Expected: 4")
		t.Fail()
	}
}

func TestExistingMidRangeElementSearch(t *testing.T) {
	zoneDataMap := CreateZoneDataMap()
	locator := &main.USPSZoneLocator{
		ZoneDataMap: zoneDataMap,
	}
	zone, err := locator.FindZone("0", 8)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if zone != 4 {
		fmt.Println("Incorrect zone returned. Zone: ", zone, " Expected: 4")
		t.Fail()
	}
}

func TestFirstRangeSearch(t *testing.T) {
	zoneDataMap := CreateZoneDataMap()
	locator := &main.USPSZoneLocator{
		ZoneDataMap: zoneDataMap,
	}
	zone, err := locator.FindZone("0", 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if zone != 1 {
		fmt.Println("Incorrect zone returned. Zone: ", zone, " Expected: 1")
		t.Fail()
	}
}

func TestLastRangeSearch(t *testing.T) {
	zoneDataMap := CreateZoneDataMap()
	locator := &main.USPSZoneLocator{
		ZoneDataMap: zoneDataMap,
	}
	zone, err := locator.FindZone("0", 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if zone != 5 {
		fmt.Println("Incorrect zone returned. Zone: ", zone, " Expected: 5")
		t.Fail()
	}
}

func TestElementNotInRangeSearch(t *testing.T) {
	zoneDataMap := CreateZoneDataMap()
	locator := &main.USPSZoneLocator{
		ZoneDataMap: zoneDataMap,
	}
	zone, err := locator.FindZone("0", 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if zone != -1 {
		fmt.Println("Incorrect zone returned. Zone: ", zone, " Expected: -1, since the element doesnt belong to any ranges.")
		t.Fail()
	}
}
