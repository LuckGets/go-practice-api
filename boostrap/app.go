package boostrap

type Application struct {
	Env *Env
}

func LoadAppConfig() Application {
	app := &Application{}
	app.Env = NewEnv()

	return *app
}
