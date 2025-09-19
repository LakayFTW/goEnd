/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/static/**/*.html", // <- deine index.html liegt hier
    "./app/**/*.go", // <- falls du Tailwind-Klassen in Go-Templates nutzt
    "./src/**/*.html", // Scan HTML-Dateien
    "./src/**/*.go",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
