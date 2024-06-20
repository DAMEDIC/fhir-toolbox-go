package rest

import (
	"time"
)

type Config struct {
	Base         string `json:"base"`
	Timezone     string `json:"timezone"`
	MaxCount     int    `json:"maxCount"`
	DefaultCount int    `json:"defaultCount"`
}

var DefaultConfig = Config{
	Timezone:     time.Now().Location().String(),
	MaxCount:     500,
	DefaultCount: 500,
}
