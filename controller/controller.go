package controller

type Controller interface {
	ExitGracefully()
	ProcessEventsFromQueue()
}
