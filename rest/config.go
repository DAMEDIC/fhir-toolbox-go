package rest

import (
	"cmp"
	"os"
)

type Config struct {
	Base string `json:"base"`
}

var DefaultConfig = Config{
	Base: cmp.Or(os.Getenv("SERVER_BASE"), ""),
}
