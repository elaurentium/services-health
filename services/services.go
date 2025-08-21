package services

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/elaurentium/services-health/cmd/banner"
)

type Service struct {
	Name   string
	Status string
}

type Services struct {
	Items  []Service
	Server string
}

func (s *Services) Health() {
	for {
		for i := range s.Items {
			serviceStatus, err := getServiceStatus(s.Items[i].Name, s.Server)
			if err != nil {
				log.Printf("error getting status of %s: %v\n", s.Items[i].Name, err)
				continue
			}

			s.Items[i].Status = serviceStatus

			banner.Usage(s.Items[i].Name, s.Items[i].Status)

			if s.Items[i].Status == "STOPPED" {
				err = startService(s.Items[i].Name, s.Server)
				if err != nil {
					log.Printf("error starting %s: %v\n", s.Items[i].Name, err)
				}
			}
		}
		time.Sleep(15 * time.Second)
	}
}

func getServiceStatus(name string, server string) (string, error) {
	cmd := exec.Command("sc", server, "query", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	outStr := strings.ToUpper(string(output))

	if strings.Contains(outStr, "RUNNING") {
		return "RUNNING", nil
	}

	if strings.Contains(outStr, "STOPPED") {
		return "STOPPED", nil
	}

	return "", fmt.Errorf("unknown status for %s: %s", name, outStr)
}

func startService(name string, server string) error {
	cmd := exec.Command("sc", server, "start", name)
	return cmd.Run()
}
