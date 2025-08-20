package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/elaurentium/services-health/cmd/banner"
)

var (
	PROD = os.Getenv("PROD_SERVER")
)

type Service struct {
	Name   string
	Status string
}

type Services struct {
	Items []Service
}

func (s *Services) Health() {
	//services := []string{"APP_JOB", "APPHARPIA_1", "APPHARPIA_2"}

	for {
		for i := range s.Items {
			cmd := exec.Command("sc", PROD, "query", s.Items[i].Name)

			output, err := cmd.CombinedOutput()
			fmt.Printf("%s\n", output)
			if err != nil {
				log.Println(err)
			}

			outStr := strings.ToUpper(string(output))

			if strings.Contains(outStr, "RUNNING") {
				s.Items[i].Status = "RUNNING"
			} else if strings.Contains(outStr, "STOPPED") {
				s.Items[i].Status = "STOPPED"

				startService := exec.Command("sc", PROD, "start", s.Items[i].Name)

				if err := startService.Run(); err != nil {
					log.Println(err)
				}
			}
			banner.Usage(s.Items[i].Name, s.Items[i].Status)
		}
		time.Sleep(15 * time.Second)
	}
}
