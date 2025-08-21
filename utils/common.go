package utils

const (
	// Version is the version of the application
	Version = "0.0.1"

	// Name is the name of the application
	Name = "services-health"

	// Description is the description of the application
	Description = "Monitor Services"
)

var (
	// .env embedded file
	Server = "\\\\10.8.0.40"

	// logger
	//Logger = func(server string) *log.Logger {
	//	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	//}
)
