package testing

import (
	"fmt"
	"github.com/ZonesProblem/server"
	"os"
	"testing"
)

func writeFile(fileLoc string, content string) (err error) {
	fh, err := os.Create(fileLoc)
	if err != nil {
		return err
	}
	if _, err := fh.WriteString(content); err != nil {
		return err
	}
	return nil
}

func TestValidCsv(t *testing.T) {
	csv := `valid,csv,text
10000,20000,1`
	writeFile("/tmp/test.txt", csv)
	reader := &main.CsvReader{
		"/tmp/test.txt",
	}
	result, err := reader.LoadData()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	expected := []main.ZipCodeZoneData{
		main.ZipCodeZoneData{
			10000,
			20000,
			1,
		},
	}
	if len(result) != len(expected) {
		fmt.Println("Length of actual and expected results doesn't match.")
		t.Fail()
	}
	for i := range result {
		fmt.Println(result[i], ":", expected[i])
		if result[i] != expected[i] {
			t.Fail()
		}
	}
	os.Remove("/tmp/test.txt")
}

func TestInvalidCsv(t *testing.T) {
	csv := `invalid:   csv:    text:
	10000:   20000:     1`
	writeFile("/tmp/test.txt", csv)
	reader := &main.CsvReader{
		"/tmp/test.txt",
	}
	_, err := reader.LoadData()
	if err == nil {
		fmt.Println("Parsed invalid csv.")
		t.FailNow()
	}
	os.Remove("/tmp/test.txt")
}

func TestValidCsvWithMoreFields(t *testing.T) {
	csv := `valid,csv,text,with,more,fields
10000,20000,1,1,1,1`
	writeFile("/tmp/test.txt", csv)
	reader := &main.CsvReader{
		"/tmp/test.txt",
	}
	_, err := reader.LoadData()
	if err == nil {
		fmt.Println("Parsed csv with more than 3 fields in the data")
		t.FailNow()
	}
	os.Remove("/tmp/test.txt")
}
