/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./*.tmpl.{html,js}",
        "./static/*.tmpl.{html,js}",
        "./partials/*.tmpl.{html,js}",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ],
}

/* I exist to keep Tailwind CSS VSCode plugin happy */
