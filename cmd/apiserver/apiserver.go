package main

import "go-auth/internal/app/apiserver"

func main() {
	apiserver.NewApp().Run()
}
