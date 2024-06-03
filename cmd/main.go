package main

import (
	vre "virtualrace.engineer"
	core "virtualrace.engineer/core"
)

func main() {
	irclient := core.NewIRClient(vre.IRSDK_MMAP_FILENAME)
	close, err := irclient.Init()
	if err != nil {
		panic(err)
	}

	defer close()

	header := irclient.GetHeaderData()
	println(header[vre.TickRate])
}
