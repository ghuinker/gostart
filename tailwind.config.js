const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  darkMode: 'media',
  content: [`./views/**/*.templ`, `./assets/**/*.{css,vue,js}`],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter var', ...defaultTheme.fontFamily.sans]
      }
    }
  },
  variants: {},
  plugins: [require('@tailwindcss/forms')]
}
