/*

esc embeds files into go programs and provides http.FileSystem interfaces
to them.

It adds all named files or files recursively under named directories at the
path specified. The output file provides an http.FileSystem interface with
zero dependencies on packages outside the standard library.

Usage:
	esc [-o outfile.go] [-pkg package] [-prefix prefix] [name ...]

The flags are:
	-o=""
		output filename, defaults to stdout
	-pkg="main"
		package name of output file, defaults to main
	-prefix=""
		strip given prefix from filenames

Accessing Embedded Files

After producing an output file, the assets may be accessed with the FS()
function, which takes a flag to use local assets instead (for local
development).

Example

Embedded assets can be served with HTTP using the http.FileServer. Assuming you have a directory 
structure similar to the following:

	/..
	/main.go
	/static/index.html
	/static/css/style.css

Where main.go contains:

	package main

	import (
		"log"
		"net/http"
	)
	
	func main() {
		//FS() is created by esc and returns a http.Filesystem compatible with http.FileServer
		http.Handle("/static/", http.FileServer(FS(false)))
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

1. Execute esc -o static.go static to generate the embedded data
2. go run main.go and static.go to start the server
3. access http://localhost:8080/static/ to view the files
*/
package main
