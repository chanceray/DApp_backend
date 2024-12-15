package main

import (
	"tce/CollectRoute"
	"tce/init"
)

func main() {
	r, db := Init.InitGinGrom()
	setup := Init.InitFabric()
	r = CollectRoute.CollectRoute(r, db, setup)
	if r.Run(":8000") != nil {
		return
	}
}
