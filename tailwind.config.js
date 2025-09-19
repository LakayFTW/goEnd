/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/static/**/*.html", // <- deine index.html liegt hier
    "./app/**/*.go", // <- falls du Tailwind-Klassen in Go-Templates nutzt
    "./src/**/*.html", // Scan HTML-Dateien
    "./src/**/*.go",
  ],
  theme: {
    extend: {
      colors: {
        text: "#f6eddc",
        background: "#0c0903",
        primary: "#e2c28a",
        secondary: "#228c5e",
        accent: "#5ab8d7",
      },
    },
  },
  plugins: [],
};
