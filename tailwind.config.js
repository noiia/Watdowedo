/** @type {import('tailwindcss').Config} */

const plugin = require("tailwindcss/plugin");

module.exports = {
	content: ["./web/**/*.{tmpl,html}"],
	theme: {
		extend: {
			screens: {
				min: "200px",
				"over-cellphone": "500px",
			},
			fontFamily: {
				display: ["Inter", "system-ui", "sans-serif"],
			},
			colors: {
				primary: "#F0C9A7",
				"secondary-orange": "#E1934F",
				"primary-content": "#000813",
				secondary: "#3d8a00",
				"secondary-content": "#010700",
				accent: "#005fff",
				"accent-content": "#cfe2ff",
				neutral: "#0f0e05",
				"neutral-content": "#c8c8c5",
				"false-white": "#fbf1e9",
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
	},
	variants: {
		extend: {
			backgroundColor: ["radio-checked"],
		},
	},
	plugins: [
		require("autoprefixer"),
		plugin(({ addVariant, e }) => {
			addVariant("radio-checked", ({ modifySelectors, separator }) => {
				modifySelectors(({ className }) => {
					const eClassName = e(`radio-checked${separator}${className}`);
					const selector = 'input[type="radio"]';
					return `${selector}:checked ~ .${eClassName}`;
				});
			});
		}),
	],
};
