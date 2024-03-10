package api

import (
	"adg/e-commerce/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

var AppConfig = &config.Configuration{}

func RegisterHandlers(r *gin.Engine) {

	os.Setenv("ENV", "local")

	env := os.Getenv("ENV")

	if !strings.EqualFold(env, "") {
		env = config.DefaultEnv
	}
	cfgFile := "../../config/app." + env + ".yaml"
	AppConfig, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		log.Fatalf("Failed to read config file. %v", err.Error())
	}

	fmt.Println(AppConfig)

}
