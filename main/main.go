package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Custom file system handler.
// Look at https://gist.github.com/hauxe/f2ea1901216177ccf9550a1b8bd59178#file-http_static_correct-go
// and https://dev.to/hauxe/golang-http-serve-static-files-correctly-2oj2
type FileSystem struct {
	//internal file system
	fs http.FileSystem
}

func (fs FileSystem) Open(path string) (http.File, error) {
	// Get file object using path
	file, error := fs.fs.Open(path)
	// return error if there is an error
	if error != nil {
		return nil, error
	}

	// get file info
	fileInfo, error := file.Stat()

	// if fileInfo is a directory
	if fileInfo.IsDir() {
		// get the path to the index.html inside the directory
		// and try to find it
		index := strings.TrimSuffix(path, "/") + "/index.html"

		if _, error := fs.fs.Open(index); error != nil {
			return nil, error
		}
	}

	return file, nil
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Main received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	log.Print("nonzz-server started.")

	// Custom FileServer for Homepage
	homeFileServer := http.FileServer(FileSystem{http.Dir("home")})

	modulesFileServer := http.FileServer(FileSystem{http.Dir("modules")})

	// Creating FileServer for html, javascript, and css
	//fs := http.FileServer(http.Dir("static"))

	// handler for domain
	http.Handle("/", homeFileServer)

	// Sets up FileServer for accessing homepage
	http.Handle("/home/", http.StripPrefix(strings.TrimRight("/home/", "/"), homeFileServer))

	// Sets up FileServer for modules
	http.Handle("/modules/", http.StripPrefix(strings.TrimRight("/modules/", "/"), modulesFileServer))

	// some handler that returns "Hello World!"
	http.HandleFunc("/main/", helloWorldHandler)

	// When the user requests with just the domain it pulls up index.html
	//http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
