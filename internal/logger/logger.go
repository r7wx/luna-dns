/*
MIT License
Copyright (c) 2022 r7wx
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
