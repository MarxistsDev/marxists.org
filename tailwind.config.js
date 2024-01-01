/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./views/*.templ'],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/typography'), require("daisyui")],
}
