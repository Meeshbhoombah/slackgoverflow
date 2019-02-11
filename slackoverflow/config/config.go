package config

import (
        "os"
	"log"
        "reflect"
)

func Load(c interface{}) error {
	config := reflect.ValueOf(c).Elem()

        for lineNo := 0; lineNo < config.NumField(); lineNo++ {
		line := config.Type().Field(lineNo)
		envVar := line.Tag.Get("env")
                val := os.Getenv(envVar)
                log.Println(val)
	}

	return nil
}
