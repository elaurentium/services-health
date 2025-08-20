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
		Server: "\\\\10.8.0.40",
		Items: []services.Service{
			{Name: "APPHARPIA"},
			//{Name: "APPHARPIA_1"},
			//{Name: "APPHARPIA_2"},
		},
	}

	s.Health()
}
