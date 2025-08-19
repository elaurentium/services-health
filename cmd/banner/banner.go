package banner

import (
	"time"

	"github.com/fatih/color"
)

const (
	banner = `
		*   ******** ******** *******   **      ** **   ******    *******    ******** *
		*  **////// /**///// /**////** /**     /**/**  **////**  **/////**  **////// *
		* /**       /**      /**   /** /**     /**/** **    //  **     //**/**        *
		* /*********/******* /*******  //**    ** /**/**       /**      /**/********* *
		* ////////**/**////  /**///**   //**  **  /**/**       /**      /**////////** *
		*        /**/**      /**  //**   //****   /**//**    ** //**     **        /**
		*  ******** /********/**   //**   //**    /** //******  //*******   ******** *
		* ////////  //////// //     //     //     //   //////    ///////   //////// *
	`
)


func PrintBanner() {
	color.New(color.FgGreen).Add(color.Bold).Println(banner)
}

func Usage(service string, status string) {
	color.New(color.FgGreen).Add(color.Bold).Printf(
		"[%s]::%s [%s]\n",
		time.Now().Format("2006-01-02 15:04"), service, status,
	)
}
