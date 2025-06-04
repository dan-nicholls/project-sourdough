/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/web/templates/*.templ"],
  theme: {
    extend: {
		colors: {
			background: "#FFF5E1",
			text: {
				light: "#333333",
				dark: "#FAFAFA",
			},
			pink: '#FF6F91',
			blue: '#4D9DE0',
			yellow: '#FFE156',
			green: '#A0E8AF',
			purple: '#8E6EC2',
			red: '#FF8C8C',
		},
		fontFamily: {
			sans: ['Outfit', 'sans-serif']
		}	
	},
  },
  plugins: [],
}

