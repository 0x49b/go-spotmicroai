package abort_controller

import (
	"github.com/lichtwellenreiter/go-spotmicroai/config"
	"sync"
)

func AbortController(wg *sync.WaitGroup, configuration config.Configuration) {
	defer wg.Done()
}

func processEventsFromQueue(){

}
