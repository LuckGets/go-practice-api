package boostrap

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	ServerPort    string `env:"SERVER_PORT"`
	DBurl         string `env:"DATABASE_URL"`
	DBName        string `env:"DATABASE_NAME"`
	DBPassword    string `env:"DATABASE_PASSWORD"`
	DBMaxOpenConn int    `env:"DATABASE_MAX_OPEN_CONN"`
	DBMaxIdleConn int    `env:"DATABASE_MAX_IDLE_CONN"`
	DBMaxIdleTime string `env:"DABASE_MAX_IDLE_TIME"`
}

func NewEnv() *Env {
	env := &Env{}
	godotenv.Load()

	// Using reflection to dynamically set struct fields
	v := reflect.ValueOf(env).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envFieldName := field.Tag.Get("env")

		data := os.Getenv(envFieldName)

		if data == "" {
			log.Fatalf("There is no data for the %s environment variable", envFieldName)
		}

		if intData, err := strconv.Atoi(data); err == nil {
			v.Field(i).SetInt(int64(intData))
			continue
		}
		fmt.Printf("This is v.Field :: %v \n", v.Field(i))

		v.Field(i).SetString(data)
	}

	return env
}
