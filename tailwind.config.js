/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./templates/*.tpl.{html,js}",
        "./static/*.tpl.{html,js}",
        "./partials/*.tpl.{html,js}",
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
