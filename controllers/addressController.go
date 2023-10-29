package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
)

type DistrictData struct {
	Districts []string `json:"districts"`
}

func GetProvinces(c *gin.Context) {
	// Read province data from JSON file in the "data" directory

	var provinceData struct {
		Provinces []string `json:"provinces"`
	}

	err := readDataFromJSON("data/provinces.json", &provinceData)
	fmt.Print(err)
	fmt.Print(provinceData.Provinces)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch provinces"})
		return
	}

	c.JSON(200, gin.H{
		"provinces": provinceData.Provinces,
	})
}

func GetDistrictsByProvince(c *gin.Context) {

	// get province name from url
	province := c.Param("province")

	var districtData struct {
		Districts []string `json:"districts"`
	}

	filePath := filepath.Join("data/districtsByProvince", province+".json")

	err := readDataFromJSON(filePath, &districtData)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch districts"})
		return
	}

	c.JSON(200, gin.H{
		"districts": districtData.Districts,
	})
}

func GetMunicipalsByDistrict(c *gin.Context) {

	// get district name from url
	district := c.Param("district")

	var municipalData struct {
		Municipals []string `json:"municipals"`
	}

	filePath := filepath.Join("data/municipalsByDistrict", district+".json")

	err := readDataFromJSON(filePath, &municipalData)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch municipals"})
		return
	}

	c.JSON(200, gin.H{
		"municipals": municipalData.Municipals,
	})
}

func readDataFromJSON(filename string, dataStructure interface{}) error {
	// Get the full path to the JSON file
	fullPath := filepath.Join(".", filename)

	// Open the file
	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the JSON data
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Parse the JSON data into the provided data structure
	if err := json.Unmarshal(data, dataStructure); err != nil {
		return err
	}

	return nil
}
