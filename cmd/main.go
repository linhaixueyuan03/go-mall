package main

import "go-mall/server"

func main() {
	r := server.NewRouter()
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
