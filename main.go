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
			САНЯ ЛОШОК
		</div>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
