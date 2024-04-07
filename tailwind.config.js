/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        'internal/templates/*.templ',
    ],
    theme: {
        colors: {
            background: "#222831",
            primary: "#00ADB5",
            "primary-accent": "#00fff5",
            "secondary": "#393E46",
            "background-foreground": "#fff",
            "primary-foreground": "#393E46",
            "secondary-foreground": "#fff"
        },
        extend: {

        },
    },
    plugins: [],
}

