package config

import (
	"fmt"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func scanEnv(message string, variable *string) {
	fmt.Printf("- %s: %s", message, console_decoration.BOLD)
	fmt.Scanf("%s", variable)
	fmt.Print(console_decoration.RESET)
}

func scanEnvWithDefaultValue(message string, variable *string, defaultValue string) {
	fmt.Printf("- %s (default: %s): %s", message, defaultValue, console_decoration.BOLD)
	fmt.Scanf("%s", variable)
	fmt.Print(console_decoration.RESET)

	if *variable == "" {
		*variable = defaultValue
	}
}
