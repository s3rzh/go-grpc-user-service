package main

import "github.com/s3rzh/go-grpc-user-service/internal/app"

const configPath = "configs"

func main() {
	app.Run(configPath)
}
