package web

func Shutdown() error {
	if app == nil {
		return nil
	}

	return app.Shutdown()
}
