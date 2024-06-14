package rest

import (
	"cmp"
	"os"
	"time"
)

type Config struct {
	Base     string `json:"base"`
	Timezone string `json:"timezone"`
}

var DefaultConfig = Config{
	Base:     cmp.Or(os.Getenv("SERVER_BASE"), ""),
	Timezone: time.Now().Location().String(),
}
