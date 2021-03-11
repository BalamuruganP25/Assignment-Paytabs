package config

import (
	INC "Paytabs/src/messages"
	"fmt"
	"os"
)

func ReadEnvironmentVariable() {

	fmt.Println("Reading Environment Variable")

	_, ApiIpStatus := os.LookupEnv("Paytaps_API_IP")
	if !ApiIpStatus {
		fmt.Println("api ip  Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Apidetails.ApiIp = os.Getenv("Paytaps_API_IP")
		fmt.Println("api ip :", INC.Apidetails.ApiIp)
	}
	_, PortStatus := os.LookupEnv("Paytaps_API_PORT")
	if !PortStatus {
		fmt.Println("api port  Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Apidetails.ApiPort = os.Getenv("Paytaps_API_PORT")
		fmt.Println("api port :", INC.Apidetails.ApiPort)
	}
	_, logfoldernamestatus := os.LookupEnv("Paytaps_LOG_FOLDER_NAME")
	if !logfoldernamestatus {
		fmt.Println("logfoldername Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Logserviceconfig.Logfoldername = os.Getenv("Paytaps_LOG_FOLDER_NAME")
		fmt.Println("logfoldername :", INC.Logserviceconfig.Logfoldername)
	}
	_, logfilepathstatus := os.LookupEnv("Paytaps_LOG_FILEPATH")
	if !logfilepathstatus {
		fmt.Println("logfilepath Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Logserviceconfig.Logfilepath = os.Getenv("Paytaps_LOG_FILEPATH")
		fmt.Println("logfilepath :", INC.Logserviceconfig.Logfilepath)
	}
}
