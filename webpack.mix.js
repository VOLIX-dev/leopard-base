let mix = require('laravel-mix');

mix.sass("./resources/sass/app.scss", "./static/css/app.css")
    .postCss("./resources/css/tailwind.css", "./static/css/tailwind.css", [
        require('tailwindcss'),
    ])
    .ts("./resources/ts/app.ts", "./static/js/app.js");