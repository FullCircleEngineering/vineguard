package environ

import (
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type Environ struct {
	TTNKey string `required:"true" envconfig:"TTN_KEY"`
}

var once sync.Once
var env Environ
var err error

func Get() (*Environ, error) {
	once.Do(func() {
		err = envconfig.Process("", &env)
	})
	return &env, err
}
