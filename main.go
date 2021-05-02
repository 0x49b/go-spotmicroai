package main

import (
	"fmt"
	"github.com/lichtwellenreiter/go-spotmicroai/config"
	abort_controller "github.com/lichtwellenreiter/go-spotmicroai/controller/abort-controller"
	"sync"
)

var configuration config.Configuration

func main() {

	configuration = config.ReadConfiguration()

	fmt.Println(configuration.AbortController.GpioPort)

	var wg sync.WaitGroup
	wg.Add(4)

	// remoteControllerChannel := make(chan string)
	// go remote_controller.PS4Controller(&wg, configuration)
	go abort_controller.AbortController(&wg, configuration)

	wg.Wait()

}
