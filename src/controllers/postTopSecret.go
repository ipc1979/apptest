package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"

	"appml/src/domain/entities"
	"appml/src/infra"
	"appml/src/services"
)

// PostTopSecret comment
func PostTopSecret(c *gin.Context) {

	var bodyRequest entities.SatellitesRequest
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		logger.Error("Save Satellites Validation Error : " + err.Error())
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	dbSatellite := infra.GetSatelliteDB()
	dbSatelliteMeasure := infra.GetSatelliteMeasureDB()

	message, err := services.SaveSatellitesMeasures(dbSatellite, dbSatelliteMeasure, bodyRequest)
	if err {
		logger.Error("Save Satellites Error : " + message)
		c.JSON(404, gin.H{"error": message})
		return
	}

	logger.Info("Save Satellites Succesfully : " + message)
	c.JSON(201, gin.H{"message": message})

}
