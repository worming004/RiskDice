package main

import (
	"strconv"
)

type configurationSglt struct {
	NbrOfSimulation int
	Mode            int
}

var instance *configurationSglt

func GetConfiguration() *configurationSglt {
	if instance == nil {
		instance = &configurationSglt{} // <--- NOT THREAD SAFE
	}
	return instance
}

type Printable interface {
	Print() string
}

func Print(conf configurationSglt) string {
	return "NumberOfSimulation : " + strconv.Itoa(conf.NbrOfSimulation)
}
