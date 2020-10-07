package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"

	"appml/src/domain/entities"
	"appml/src/infra"
	"appml/src/services"
)

// PostTopSecretSplit comment
func PostTopSecretSplit(c *gin.Context) {

	satelliteName := c.Params.ByName("satellite_name")
	var bodyRequest entities.MeasureRequest
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		logger.Error("Save Satellite Validation Error :" + err.Error())
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	dbSatellite := infra.GetSatelliteDB()
	dbSatelliteMeasure := infra.GetSatelliteMeasureDB()

	message, err := services.SaveSatellitesMeasures(dbSatellite, dbSatelliteMeasure,
		entities.SatellitesRequest{
			Satellites: []entities.SatelliteRequest{
				{
					Name:     satelliteName,
					Distance: bodyRequest.Distance,
					Message:  bodyRequest.Message}}})
	if err {
		logger.Error("Save Satellite Error : " + message)
		c.JSON(404, gin.H{"error": message})
		return
	}

	logger.Info("Save Satellite Succesfully : " + message)
	c.JSON(201, gin.H{"message": message})

}
