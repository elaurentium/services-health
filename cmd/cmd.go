package main

import (
	"github.com/elaurentium/services-health/cmd/banner"
	"github.com/elaurentium/services-health/services"
)

func init() {
	banner.PrintBanner()
}

func main() {
	s := &services.Services{
		Items: []services.Service{
			{Name: "XblGameSave"},
			{Name: "APPHARPIA_1"},
			{Name: "APPHARPIA_2"},
		},
	}

	s.Health()
}
