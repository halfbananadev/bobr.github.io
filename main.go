package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html>
	<head>
 		<link rel="stylesheet" href="styles.css">
		<style>
		body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            font-family: Sora, Arial, sans-serif;
            background-color: #10203A;
            color: white;
        }
        .App {
            font-size: 2em; /* Adjust this value to increase or decrease the font size */
        }
		</style>
	</head>
	<body>
		<div class="App">
  		<div id="shadowBox">
    			<h3 class="rainbow rainbow_text_animated"><a target="_blank" href="https://vk.com/skarushkin">САНЯ ЛОШОК</a></h3>
		</div>
		</div>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
