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
			{Name: "APPHARPIA"},
			//{Name: "APPHARPIA_1"},
			//{Name: "APPHARPIA_2"},
		},
	}

	s.Health()
}
