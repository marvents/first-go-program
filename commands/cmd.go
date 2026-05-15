package commands

import (
	"fmt"
	"program/services"
)

func start() {
	count := 1
	choice := 0
	fmt.Print("============\nchoose a service\n")
	for key := range services.Services {
		fmt.Printf("%d: %s\n", count, key)
		count++
	}

	fmt.Scanln(&choice)

	switch choice { // stoping and starting program again is counts as restart, that's why I added 1 on restartCount ;)
	case 1:
		services.Services["database"]["status"] = 1
		services.Services["database"]["restartCount"]++
	case 2:
		services.Services["nginx"]["status"] = 1
		services.Services["nginx"]["restartCount"]++
	case 3:
		services.Services["frontend"]["status"] = 1
		services.Services["frontend"]["restartCount"]++
	case 4:
		services.Services["backend"]["status"] = 1
		services.Services["backend"]["restartCount"]++
	}

	fmt.Print("service started successfully\n")

}

func stop() {
	count := 1
	choice := 0
	fmt.Print("============\nchoose a service\n")
	for key := range services.Services {
		fmt.Printf("%d: %s\n", count, key)
		count++
	}

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		services.Services["database"]["status"] = 0
	case 2:
		services.Services["nginx"]["status"] = 0
	case 3:
		services.Services["frontend"]["status"] = 0
	case 4:
		services.Services["backend"]["status"] = 0
	}

	fmt.Print("service stopped successfully\n")
}

func restart() { // it's basically like start :), cause those services it's fake now.
	count := 1
	choice := 0
	fmt.Print("============\nchoose a service\n")
	for key := range services.Services {
		fmt.Printf("%d: %s\n", count, key)
		count++
	}

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		services.Services["database"]["status"] = 1
		services.Services["database"]["restartCount"]++
	case 2:
		services.Services["nginx"]["status"] = 1
		services.Services["nginx"]["restartCount"]++
	case 3:
		services.Services["frontend"]["status"] = 1
		services.Services["frontend"]["restartCount"]++
	case 4:
		services.Services["backend"]["status"] = 1
		services.Services["backend"]["restartCount"]++
	}
	
	fmt.Print("service restarted successfully\n")
}

func status() {
	count := 1
	choice := 0
	fmt.Print("============\nchoose a service\n")
	for key := range services.Services {
		fmt.Printf("%d: %s\n", count, key)
		count++
	}

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Printf("service: database | running: %t\n", services.Services["database"]["status"] == 1)
	case 2:
		fmt.Printf("service: nginx | running: %t\n", services.Services["nginx"]["status"] == 1)
	case 3:
		fmt.Printf("service: frontend | running: %t\n", services.Services["frontend"]["status"] == 1)
	case 4:
		fmt.Printf("service: backend | running: %t\n", services.Services["backend"]["status"] == 1)
	}
}

func Cmd(cmd string) {
	switch cmd {
	case "start":
		start()
	case "stop":
		stop()
	case "restart":
		restart()
	case "status":
		status()
	case "help":
		fmt.Print("\n======================\nCommands Guide:\nstart: start a service\nstop: stop a service\nrestart: restart a service\nstatus: check the status of a service\nhelp: show this message\n")
	default:
		fmt.Print("\n======================\nCommands Guide:\nstart: start a service\nstop: stop a service\nrestart: restart a service\nstatus: check the status of a service\nhelp: show this message\n")
	}
}
