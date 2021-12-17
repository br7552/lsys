package main

import (
	"flag"
	"log"
	"net/http"
)

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
</head>
<body>
	<form id="fractalForm">
		<label for="axiom">Axiom:</label></br>
		<input type="text" id="axiom" value="F++F++F"><br>
		<label for="rules">Rules:</label></br>
		<input type="text" id="rules" value="F=F-F++F-F"><br>
		<label for="depth">Depth:</label><br>
		<input type="number" id="depth" value="3"><br>
		<label for="angle">Angle:</label><br>
		<input type="number" id="angle" value="60"><br>
		<label for="step">Step:</label><br>
		<input type="number" id="step" value="3"><br>
		<input type="button" onclick="generate()" value="Generate">
	</form>


	<div id="output"></div>

	<script>
		function generate() {
			var form = document.getElementById("fractalForm");
			var axiom = form["axiom"].value;
			var rules = form["rules"].value.split(" ").map(function(rule) {
				return rule.split("=");
			});
			var depth = parseInt(form["depth"].value);
			var angle = parseFloat(form["angle"].value);
			var step = parseInt(form["step"].value);

			fetch("http://localhost:4000/v1/fractals", {
				method: "POST",
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					axiom: axiom,
					rules: Object.fromEntries(rules),
					angle: angle,
					height: 66,
					width: 100,
					step: step,
					depth: depth
				})
			}).then(
				function (response) {
					response.json().then(function(obj) {
						var fractal = obj.fractal.Data
						document.getElementById("output").innerHTML =
							'<pre>' + fractal + '</pre>';
					});
				},
				function(err) {
					document.getElementById("output").innerHTML = err;
				}
			);
		}
	</script>
</body>
</html>`

func main() {
	addr := flag.String("addr", ":9000", "Server Address")

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}))
	log.Fatal(err)
}
