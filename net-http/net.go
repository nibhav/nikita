package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type header to HTML
		w.Header().Set("Content-Type", "text/html")
		// Write the HTML response to the response writer
		fmt.Fprint(w, `
            <!DOCTYPE html>
            <html>
                <head>
                    <title>Hello, world!</title>
                </head>
                <body>
                    <h1>Hello, world!</h1>
					<h2>HTTP server Program in Go Programming</h2>
					<h3>Trying to run Program on server</h3>
                </body>
            </html>
        `)
	})
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "<h1><i>Hello, world!</i></h1>")
	//})
	http.ListenAndServe(":8080", nil)
} //
