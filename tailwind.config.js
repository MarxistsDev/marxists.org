/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./views/*.gohtml'],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/typography'), require("daisyui")],
}

//npx tailwindcss -o ./www/styles/style.css --watch 
