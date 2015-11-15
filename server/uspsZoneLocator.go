package main

import (
	"fmt"
	"math"
)

type USPSZoneLocator struct {
	ZoneDataMap map[string][]ZipCodeZoneData
}

func (locator USPSZoneLocator) FindZone(origCode string, destCode int) (zone int, err error) {
	Info.Println("Locating stuff")
	var found bool
	zipZoneData := locator.ZoneDataMap[origCode]
	var low int = 0
	var high int = len(zipZoneData) - 1
	var lastIndex int
	for !found && low <= high {
		fmt.Println("Low:", low, "High:", high)
		mid := float64(low+high) / float64(2)
		if mid == float64(int64(mid)) {
			lastIndex = int(mid)
			locator.updateCursor(origCode, destCode, lastIndex, &low, &high, &found)
		} else {
			lastIndex = int(math.Floor(mid))
			locator.updateCursor(origCode, destCode, lastIndex, &low, &high, &found)
			if found {
				continue
			}
			lastIndex = int(math.Ceil(mid))
			locator.updateCursor(origCode, destCode, lastIndex, &low, &high, &found)
		}
	}
	if found {
		fmt.Println("Zone:", locator.ZoneDataMap[origCode][lastIndex].Zone, "Low:",
			locator.ZoneDataMap[origCode][lastIndex].StartCode, "High:", locator.ZoneDataMap[origCode][lastIndex].EndCode)
		return locator.ZoneDataMap[origCode][lastIndex].Zone, nil
	}
	return -1, nil
}

func (locator *USPSZoneLocator) updateCursor(origCode string, destCode int, index int, low *int, high *int, found *bool) {
	zipZoneData := locator.ZoneDataMap[origCode]
	switch {
	case destCode >= zipZoneData[index].StartCode && destCode <= zipZoneData[index].EndCode:
		*found = true

	case destCode < zipZoneData[index].StartCode:
		*high = index - 1

	case destCode > zipZoneData[index].EndCode:
		*low = index + 1
	}
}
