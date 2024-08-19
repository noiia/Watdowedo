/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./src/**/*.{tmpl,html,js}"],
	theme: {
		fontFamily: {
			display: ["Inter", "system-ui", "sans-serif"],
		},
		colors: {
			primary: "#0094eb",
			"primary-content": "#000813",
			secondary: "#3d8a00",
			"secondary-content": "#010700",
			accent: "#005fff",
			"accent-content": "#cfe2ff",
			neutral: "#0f0e05",
			"neutral-content": "#c8c8c5",
			"base-100": "#fffdff",
			"base-200": "#dedcde",
			"base-300": "#bebcbe",
			"base-content": "#161616",
			info: "#2e93ff",
			"info-content": "#010816",
			success: "#008200",
			"success-content": "#d3e6d1",
			warning: "#df7400",
			"warning-content": "#120500",
			error: "#be123c",
			"error-content": "#f43f5e",
		},
	},
	plugins: [require("autoprefixer")],
};
