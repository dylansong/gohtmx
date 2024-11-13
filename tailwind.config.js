
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["views/**/*.templ", "views/**/*.go"],
  theme: {
    container: {
      center: true,
      padding: {
        DEFAULT: "1rem",
        mobile: "2rem",
        tablet: "4rem",
        desktop: "5rem",
      },
    } 
  },
  plugins: [],
};