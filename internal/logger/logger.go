package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

// DebugEnabled - Enable or disable debug level in logs
var DebugEnabled bool = false

// Info - Print an info log message
func Info(message any) {
	timestamp := time.Now().Format(time.RFC822)
	fmt.Printf("%v%v %s\n",
		color.MagentaString("["+timestamp+"]"),
		color.MagentaString("[luna-dns]"),
		message,
	)
}

// Debug - Print an debug log message if debug is enabled
func Debug(message any) {
	if !DebugEnabled {
		return
	}

	timestamp := time.Now().Format(time.RFC822)
	fmt.Printf("%v%v %s\n",
		color.YellowString("["+timestamp+"]"),
		color.YellowString("[luna-dns]"),
		message,
	)
}

// Error - Print an error log message
func Error(message any) {
	timestamp := time.Now().Format(time.RFC822)
	fmt.Printf("%v%v Error: %s\n",
		color.RedString("["+timestamp+"]"),
		color.RedString("[luna-dns]"),
		message,
	)
}

// Fatal - Print an error log message and quit
func Fatal(message any) {
	Error(message)
	os.Exit(-1)
}
