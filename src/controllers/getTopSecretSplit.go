package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"

	"appml/src/infra"
	"appml/src/services"
)

// GetTopSecret comment
func GetTopSecretSplit(c *gin.Context) {

	dbSatellite := infra.GetSatelliteDB()
	dbSatelliteMeasure := infra.GetSatelliteMeasureDB()

	position, err := services.CalculatePosition(dbSatellite, dbSatelliteMeasure)
	if err {
		logger.Error("Calculate Position Error")
		c.JSON(404, gin.H{"error": "position cant be recovered"})
		return
	}

	message, err := services.RecoverMessage(dbSatelliteMeasure)
	if err {
		logger.Error("Recover Message Error")
		c.JSON(404, gin.H{"error": "message cant be recovered"})
		return
	}

	logger.Info("Position and Message Recovered Succesfully")
	c.JSON(200, gin.H{"position": position, "message": message})

}
