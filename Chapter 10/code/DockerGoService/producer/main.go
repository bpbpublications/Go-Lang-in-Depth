package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"rest_api_mq/producer/config"
	"rest_api_mq/producer/utils"
)

func init() {

	mode := utils.GetEnvVar("GIN_MODE")
	gin.SetMode(mode)
}

func main() {

	appGin := config.CreateApp()

	addrGin := utils.GetEnvVar("GIN_ADDR")
	portGin := utils.GetEnvVar("GIN_PORT")

	log.Info().Msgf("App is up at http//:%s:%s", addrGin, portGin)
	if error := appGin.Run(fmt.Sprintf("%s:%s", addrGn, portGin)); error != nil {
		log.Fatal().Err(error).Msg("Http Server setup failed")
	}
}
