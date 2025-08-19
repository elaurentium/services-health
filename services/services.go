package services

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/elaurentium/services-health/cmd/banner"
)

var (
	status string
)

func Health() {
	//services := []string{"APP_JOB", "APPHARPIA_1", "APPHARPIA_2"}

	services := []string{"XblGameSave", "APPHARPIA_1", "APPHARPIA_2"}

	for {
		for _, serviceName := range services {
			cmd := exec.Command("sc", "query", serviceName)

			output, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}

			outStr := strings.ToUpper(string(output))

			if strings.Contains(outStr, "RUNNING") {
				status = "RUNNING"
			} else if strings.Contains(outStr, "STOPPED") {
				status = "STOPPED"

				startService := exec.Command("sudo", "sc", "start", serviceName)
				
				if err := startService.Run(); err != nil {
					log.Println("deu ruim",err)
				}
			}

			banner.Usage(serviceName, status)
		}
		time.Sleep(30 * time.Second)
	}
}

