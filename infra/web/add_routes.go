package web

func (w *WebServer) AddRoutes() {
	w.App.Get("/", w.APIS.HandlerHealth())

	w.App.Get("/current", w.APIS.HandlerGetCurrentTarget())
	w.App.Get("/targets", w.APIS.HandlerGetTargets())
	w.App.Post("/target", w.APIS.HandlerGetTargetConfiguration())

	w.App.Post("/values", w.APIS.HandlerGetVariableValues())
	w.App.Post("/delete", w.APIS.HandlerDeleteTargetConfiguration())
}
