package main

import (
	vre "virtualrace.engineer"
	core "virtualrace.engineer/core"
)

func main() {
	irclient := core.NewIRClient(vre.IRSDK_MMAP_FILENAME)
	err := irclient.Init()
	if err != nil {
		panic(err)
	}

	data := irclient.GetHeaderData()
	for _, val := range data {
		println(val)
	}
}
