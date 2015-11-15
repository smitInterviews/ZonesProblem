package main

type ZoneDataReader interface {
	LoadData() ([]ZipCodeZoneData, error)
}

type ZoneLocator interface {
	FindZone(string, int) (int, error)
}
