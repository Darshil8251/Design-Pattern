package adapter

import (
	"design-pattern/structural/adapter/logger"
	"fmt"
)

func main() {
	logger, err := logger.NewCustomLogger("development")
	if err != nil {
		fmt.Printf("failed to initialize logger: %v\n", err)
	}
	logger.Info("welcome to home")
}
