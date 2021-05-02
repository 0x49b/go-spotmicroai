package remote_controller

import (
	"context"
	"fmt"
	"github.com/lichtwellenreiter/go-spotmicroai/config"
	"github.com/mrasband/ps4"
	"log"
	"sync"
)

func PS4Controller(wg *sync.WaitGroup, configuration config.Configuration) {
	defer wg.Done()

	var inputs, err = ps4.Discover()
	if err != nil {
		fmt.Println("Cannot connect to Remote Controller")
		wg.Done()
	}

	var device *ps4.Input
	for _, input := range inputs {
		if input.Type == ps4.Controller {
			device = input
			break
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	events, _ := ps4.Watch(ctx, device)
	for e := range events {
		log.Printf("%+v\n", e)
	}
}
