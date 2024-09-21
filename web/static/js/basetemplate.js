document.addEventListener("DOMContentLoaded", function () {
	const burger = document.getElementById("navbar-burger");
	const menu = document.getElementById("navbar-menu");

	function handleResize() {
		if (window.innerWidth >= 1024) {
			menu.classList.remove("hidden");
			burger.classList.add("hidden");
		} else {
			burger.classList.remove("hidden");
			menu.classList.add("hidden");
		}
	}

	handleResize();

	window.addEventListener("resize", handleResize);
});
