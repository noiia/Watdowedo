function getRadioButtonValue(radioName) {
	return document.querySelector('input[name="' + radioName + '"]:checked').id;
}

function postFormData() {
	let jsonObject = {};
	let radioButtonName = ["walking-level", "validity", "drive"];
	radioButtonName.forEach(
		(radioName) => (jsonObject[radioName] = getRadioButtonValue(radioName))
	);
	console.log(jsonObject);

	fetch("http://www.watdowedo.local/tripbuilder/form", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(jsonObject),
	})
		.then((response) => response.text())
		.then((data) => {
			console.log("Réponse du serveur :", data);
		})
		.catch((error) => {
			console.error("Erreur lors de l'envoi du formulaire :", error);
		});
}
