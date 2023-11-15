/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./views/*.gohtml'],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
}