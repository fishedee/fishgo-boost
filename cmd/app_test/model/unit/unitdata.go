package unit

import (
	"time"
)

type Unit struct {
	UnitId     int `sqlf:"autoincr"`
	Name       string
	CreateTime time.Time `sqlf:"created"`
	ModifyTime time.Time `sqlf:"updated"`
}

type Units struct {
	Count int
	Data  []Unit
}
