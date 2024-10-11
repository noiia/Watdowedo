function getRadioButtonValue(radioName) {
	return document.querySelector('input[name="' + radioName + '"]:checked').id;
}

function postFormData() {
	let jsonObject = {};
	let destination = "destination";
	let radioButtonName = ["walking-level", "validity", "drive"];

	jsonObject[destination] = document.getElementById(destination).value;
	radioButtonName.forEach(
		(radioName) => (jsonObject[radioName] = getRadioButtonValue(radioName))
	);
	console.log(jsonObject);

	fetch("http://localhost:8080/tripbuilder/form", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(jsonObject),
	})
		.then((response) => response.text())
		.catch((error) => {
			console.error("Erreur lors de l'envoi du formulaire :", error);
		});
}
