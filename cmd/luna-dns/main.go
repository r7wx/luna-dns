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

package main

import (
	"os"

	"github.com/r7wx/luna-dns/internal/config"
	"github.com/r7wx/luna-dns/internal/engine"
	"github.com/r7wx/luna-dns/internal/logger"
	"github.com/r7wx/luna-dns/internal/logo"
)

func main() {
	logo.Print()

	args := os.Args[1:]
	if len(args) <= 0 {
		logger.Fatal("no configuration file provided")
	}
	config, err := config.Load(args[0])
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Configuration file loaded: " + args[0])

	engine, err := engine.NewEngine(config)
	if err != nil {
		logger.Fatal(err)
	}
	if err := engine.Start(); err != nil {
		logger.Fatal(err)
	}
}
