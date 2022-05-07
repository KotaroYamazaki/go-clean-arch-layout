package env

import "github.com/kelseyhightower/envconfig"

type Env struct {
	Port string `envconfig:"PORT" default:":3000"`
	Env  string `envconfig:"ENV" default:"local"`
}

var ENV Env

func Init() {
	_ = envconfig.Process("", &ENV)
}
