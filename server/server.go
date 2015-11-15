package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func NewZoneDataReader(loc string) (reader ZoneDataReader, err error) {
	newReader := &CsvReader{
		FileLoc: loc,
	}
	return newReader, nil
}

func NewZoneLocator(zoneDataMap map[string][]ZipCodeZoneData) (locator ZoneLocator, err error) {
	newZoneLocator := USPSZoneLocator{
		ZoneDataMap: zoneDataMap,
	}
	return newZoneLocator, nil
}

func GetZone(ctx *gin.Context) {
	origCode := ctx.Param("origin_code")
	destCode := ctx.Query("dest_code")
	if len(destCode) != 5 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Destination Code. Needs to be 5 digits."})
		return
	}
	Info.Println("Request origin Code:", origCode)
	if origCode != "90040" {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Origin Code Not Supported."})
		return
	}
	destCode = destCode[:3]
	destCodeInt, err := strconv.Atoi(destCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": fmt.Sprintf("Invalid Destination Code: %s", ctx.Query("dest_code"))})
		return
	}
	zone, err := ZipZoneLocator.FindZone(origCode, destCodeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": fmt.Sprintf("Locating zone encountered error: %s", err)})
		return
	}
	if zone == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Destination Code not found in distance matrix."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Zone " + strconv.Itoa(zone)})
}

func init() {
	logFile, err := os.OpenFile("/tmp/zones_server_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
		panic(err)
	}
	multi := io.MultiWriter(logFile, os.Stdout)
	Info = log.New(multi,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(multi,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(multi,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	//os.Setenv("GOPATH", "~/gowork2")
}

func instantiateDependencies() {
	dataReader, err := NewZoneDataReader(os.Getenv("GOPATH") + "/src/github.com/ZonesProblem/server/csvData/los-angeles-zone-data.csv")
	if err != nil {
		Error.Printf("Error instantiating zone data reader: %s\n", err)
		panic(err)
	}
	zipZoneList, err := dataReader.LoadData()
	if err != nil {
		Error.Printf("Error loading data: %s\n", err)
		panic(err)
	}
	for i := range zipZoneList {
		Info.Printf("%v \n", zipZoneList[i])
	}
	var zipZoneMatrix map[string][]ZipCodeZoneData = make(map[string][]ZipCodeZoneData)
	zipZoneMatrix["90040"] = zipZoneList
	ZipZoneLocator, err = NewZoneLocator(zipZoneMatrix)
	if err != nil {
		Error.Printf("Error instantiating zone locator: %s\n", err)
		panic(err)
	}

}

func main() {
	instantiateDependencies()
	r := gin.Default()
	r.GET("/zone/:origin_code/distancematrix", GetZone)
	r.Run(fmt.Sprintf("%s:%d", "localhost", 8080))
}
