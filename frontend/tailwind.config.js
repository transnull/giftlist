/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        brand: {
          50:  '#fdf8f6',
          100: '#f9f0eb',
          200: '#f0d9cc',
          300: '#e4b89e',
          400: '#d4896a',
          500: '#c4623e',
          600: '#b04d2e',
          700: '#923d27',
          800: '#763226',
          900: '#622c24',
        },
      },
      fontFamily: {
        sans: ['-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'Roboto', '"Helvetica Neue"', 'Arial', '"Noto Sans SC"', 'sans-serif'],
      },
    },
  },
  plugins: [],
}
