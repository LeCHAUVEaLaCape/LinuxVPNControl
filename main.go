package main

import (
	"fmt"
	"os"

	"LinuxVPNControl/pkg/vpnmanager"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : myvpnmanager <start|stop|status>")
		return
	}
	switch os.Args[1] {
	case "start":
		err := vpnmanager.StartVPN()
		if err != nil {
			fmt.Println("Error Starting VPN : ", err)
		} else {
			fmt.Println("VPN started succesfully.")
		}
	case "stop":
		err := vpnmanager.StopVPN()
		if err != nil {
			fmt.Println("Error Stoping VPN : ", err)
		} else {
			fmt.Println("VPN stopped succesfully.")
		}
	case "status":
		status, err := vpnmanager.GetStatus()
		if err != nil {
			fmt.Println("Error fetching VPN status : ", err)
		} else {
			fmt.Println("VPN Status :", status)
		}
	default:
		fmt.Println("Unknown command. Use start, stop or status.")
	}
}
