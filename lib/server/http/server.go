package http

type Server interface {
	Serve(port int)
	Stop()
}

// TODO: serve the html page for configuring on/off schedules
// TODO: api endpoints
