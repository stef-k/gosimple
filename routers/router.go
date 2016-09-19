package routers

// Initialize all routes
func init() {
    SetupBasicRoutes()
	SetupWebsocketRoutes()
	SetupV1Api()
}
