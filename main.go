package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/spf13/viper"

	"appml/src/controllers"
	"appml/src/domain/entities"
	"appml/src/infra"
)

func main() {

	viper.SetConfigName("default")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("Error reading configuration.")
		os.Exit(1)
	}

	host := viper.GetString("http.host")
	port := viper.GetString("http.port")

	logger.Info("Initializing appml ...")

	dbSatellite := infra.GetSatelliteDB()

	dbSatellite.Save(entities.Satellite{Name: "kenobi", Position: entities.Position{X: -500.0, Y: -200.0}})
	dbSatellite.Save(entities.Satellite{Name: "skywalker", Position: entities.Position{X: 100.0, Y: -100.0}})
	dbSatellite.Save(entities.Satellite{Name: "sato", Position: entities.Position{X: 500.0, Y: 100.0}})

	r := gin.Default()

	r.GET("/topsecret", controllers.GetTopSecret)

	r.POST("/topsecret", controllers.PostTopSecret)

	r.GET("/topsecret_split", controllers.GetTopSecretSplit)

	r.POST("/topsecret_split/:satellite_name", controllers.PostTopSecretSplit)

	r.Run(host + ":" + port)

}
