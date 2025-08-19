package main

import (
	"github.com/elaurentium/services-health/cmd/banner"
	"github.com/elaurentium/services-health/services"
)

func init() {
	banner.PrintBanner()
}

func main() {
	services.Health()
}