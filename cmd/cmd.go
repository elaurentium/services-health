package cmd

import (
	"github.com/elaurentium/services-health/cmd/banner"
	"github.com/elaurentium/services-health/services"
	"github.com/elaurentium/services-health/utils"
	"github.com/joho/godotenv"
)

func init() {
	banner.PrintBanner()
	godotenv.Load()
}

func Run() {
	s := &services.Services{
		Server: utils.Server,
		Items: []services.Service{
			{Name: "APPJOB"},
			{Name: "APPONCA_1"},
			{Name: "APPONCA_2"},
			{Name: "APPWEB"},
		},
	}

	s.Health()
}
