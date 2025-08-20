package main

import (
	"os"

	"github.com/elaurentium/services-health/cmd/banner"
	"github.com/elaurentium/services-health/services"
	"github.com/joho/godotenv"
)

func init() {
	banner.PrintBanner()
	godotenv.Load()
}

func main() {
	s := &services.Services{
		Server: os.Getenv("PROD_SERVER"),
		Items: []services.Service{
			{Name: "APPHARPIA"},
			//{Name: "APPHARPIA_1"},
			//{Name: "APPHARPIA_2"},
		},
	}

	s.Health()
}
