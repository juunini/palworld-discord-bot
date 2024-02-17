package config

import (
	"bufio"
	"fmt"
	"os"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func scanEnv(message string, variable *string) {
	fmt.Printf("- %s: %s", message, console_decoration.BOLD)
	*variable = scan()
	fmt.Print(console_decoration.RESET)
}

func scanEnvWithDefaultValue(message string, variable *string, defaultValue string) {
	fmt.Printf("- %s (default: %s): %s", message, defaultValue, console_decoration.BOLD)
	*variable = scan()
	fmt.Print(console_decoration.RESET)

	if *variable == "" {
		*variable = defaultValue
	}
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
