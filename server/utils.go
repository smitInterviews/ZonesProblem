package main

import (
	"strconv"
)

func convert(orig string, dest string, zone string) (int, int, int, error) {
	o, err := strconv.Atoi(orig)
	if err != nil {
		return -1, -1, -1, err
	}

	i, err := strconv.Atoi(dest)
	if err != nil {
		return -1, -1, -1, err
	}

	z, err := strconv.Atoi(zone)
	if err != nil {
		return -1, -1, -1, err
	}

	return o, i, z, nil
}
