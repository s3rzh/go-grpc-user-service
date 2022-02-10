package main

import "github.com/s3rzh/go-grpc-user-service/internal/logger"

const configPath = "configs"

func main() {
	logger.Run(configPath)
}
