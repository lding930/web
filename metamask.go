package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>MetaMask Connection</title>
			<script>
				document.addEventListener('DOMContentLoaded', function () {
					const connectButton = document.getElementById('connectButton');
					const addressDisplay = document.getElementById('addressDisplay');

					connectButton.addEventListener('click', async () => {
						try {
							const accounts = await ethereum.enable();
							const connectedAddress = accounts[0];

							const addressElement = document.createElement('p');
							addressElement.textContent = 'Address: ' + connectedAddress;
							addressDisplay.appendChild(addressElement);

							// You can also send the connected address to the server if needed
							fetch('/address', {
								method: 'POST',
								headers: {
									'Content-Type': 'application/json',
								},
								body: JSON.stringify({ address: connectedAddress }),
							});
						} catch (error) {
							console.error('Error connecting with MetaMask:', error);
						}
					});
				});
			</script>
		</head>
		<body>
			<h1>MetaMask Connection</h1>
			<button id="connectButton">Connect</button>
			<div id="addressDisplay"></div>
		</body>
		</html>
		`
		w.Write([]byte(html))
	})

	http.HandleFunc("/address", func(w http.ResponseWriter, r *http.Request) {
		// Handle the address sent from the client (if needed)
		// This is just an example; you can perform any desired server-side logic here
		fmt.Println("Received address:", r.FormValue("address"))
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

