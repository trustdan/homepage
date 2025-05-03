package main

import (
	"log"
	"net/http"

	"github.com/yourusername/particles-go/particles" // Replace with your actual import path
)

func main() {
	// Create a new particles handler for Hugo
	// The first parameter is the endpoint URL for configurations
	// The second parameter is the path to the particles.js file within your static directory
	particlesHandler := particles.NewHugoHandler(
		"/api/particles-config",
		"/js/particles.min.js",
	)

	// Register the handler to serve particle configs
	http.Handle("/api/particles-config", particlesHandler)

	// Serve static files from the 'static' directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/js/", fs)

	// Add more routes for your Hugo API here if needed

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// How to use with Hugo:
//
// 1. Add the particles.js library to your static directory:
//    - Copy particles.min.js to static/js/
//
// 2. Create the Hugo shortcode in layouts/shortcodes/particles.html
//
// 3. Build and run this Go program alongside your Hugo site
//
// 4. Use the shortcode in your Hugo content:
//    {{< particles >}} - Uses default config
//    {{< particles preset="snow" >}} - Uses a snow preset
//    {{< particles color="#ff0000" number="150" >}} - Customized
//
// Note: For production, you may want to integrate this with your
// main Hugo site server using Go plugins or embedding Hugo as a library.
