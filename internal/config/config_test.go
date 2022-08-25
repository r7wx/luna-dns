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

package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	writeFile := func(t *testing.T, content []byte) string {
		tmpFile, err := ioutil.TempFile(".", "luna-dns_test")
		if err != nil {
			t.Fatal(err)
		}

		_, err = tmpFile.Write(content)
		if err != nil {
			t.Fatal(err)
		}

		return tmpFile.Name()
	}

	filePath := writeFile(t, []byte(`addr: 0.0.0.0:5353
network: tcp
debug: true
dns:
  - addr: "8.8.8.8:53"
    network: "udp"
hosts:
  - host: google.com
    ip: 127.0.0.1
  - host: "*.test.com"
    ip: 127.0.0.1`))
	defer os.Remove(filePath)

	config, err := Load(filePath)
	if err != nil {
		t.Fatal(err)
	}

	if config.Addr != "0.0.0.0:5353" {
		t.Fatal()
	}
	if config.Network != "tcp" {
		t.Fatal()
	}
	if !reflect.DeepEqual(config.DNS, []DNS{
		{
			Addr:    "8.8.8.8:53",
			Network: "udp",
		},
	}) {
		t.Fatal()
	}
	if !reflect.DeepEqual(config.Hosts, []Host{
		{
			Host: "google.com",
			IP:   "127.0.0.1",
		},
		{
			Host: "*.test.com",
			IP:   "127.0.0.1",
		},
	}) {
		t.Fatal()
	}

	config, err = Load("ne/not_existent_path_i_hope")
	if err == nil {
		t.Fatal()
	}
}
