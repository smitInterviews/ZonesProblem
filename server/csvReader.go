package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type CsvReader struct {
	FileLoc string
}

func (c CsvReader) LoadData() (zipZoneList []ZipCodeZoneData, err error) {
	fh, err := os.Open(c.FileLoc)
	if err != nil {
		return nil, err
	}
	bufferedReader := bufio.NewReader(fh)
	reader := csv.NewReader(bufferedReader)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	for i := range records {
		if i == 0 {
			continue
		}
		if len(records[i]) != 3 {
			return nil,
				errors.New(
					fmt.Sprintf("Unexpected number of fields while parsing row."+
						"Number of fields: %d, Row: %v", len(records[i]), records[i]))
		}
		origCode, destCode, zone, err := convert(records[i][0], records[i][1], records[i][2])
		if err != nil {
			return nil, errors.New(
				fmt.Sprintf("Integer parsing issue for csv."+
					"Fields: %d, %d, %d Err: %s", records[i][0], records[i][1], records[i][2], err))
		}
		zipZoneList = append(zipZoneList,
			ZipCodeZoneData{
				StartCode: origCode,
				EndCode:   destCode,
				Zone:      zone,
			})
	}
	return zipZoneList, nil
}
