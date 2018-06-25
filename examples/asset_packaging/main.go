// Copyright 2016 José Santos <henrique_1609@me.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +Build ignore
//go:generate go run assets/generate.go
package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/maxime1907/jet"
	"github.com/maxime1907/jet/examples/asset_packaging/assets/templates"
	"github.com/maxime1907/jet/loaders/httpfs"
	"github.com/maxime1907/jet/loaders/multi"
)

// Initialize the set with both local files as well as the packaged
// views generated with `go generate` during the build step.
var views = jet.NewHTMLSetLoader(multi.NewLoader(
	jet.NewOSFileSystemLoader("./views"),
	httpfs.NewLoader(templates.Assets),
))

func main() {
	// remove in production
	views.SetDevelopmentMode(true)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view, err := views.GetTemplate("index.jet")
		if err != nil {
			log.Println("Unexpected template err:", err.Error())
		}
		view.Execute(w, nil, nil)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":9090"
	} else if !strings.HasPrefix(":", port) {
		port = ":" + port
	}

	log.Println("Serving on " + port)
	http.ListenAndServe(port, nil)
}
