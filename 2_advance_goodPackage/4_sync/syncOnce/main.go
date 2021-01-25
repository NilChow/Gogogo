package main

import (
	"Gogogo/2_advance_goodPackage/4_sync/syncOnce/userA"
	"Gogogo/2_advance_goodPackage/4_sync/syncOnce/userB"
	"Gogogo/2_advance_goodPackage/4_sync/syncOnce/worker"
)

func main() {
	userA.Init()
	userB.Init()
	worker.Init()
}