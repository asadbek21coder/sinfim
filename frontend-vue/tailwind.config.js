/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
        display: ['Manrope', 'Inter', 'system-ui', 'sans-serif'],
      },
      colors: {
        primary:                    '#041632',
        'primary-container':        '#dce7f6',
        'on-primary':               '#ffffff',
        'on-primary-container':     '#041632',
        secondary:                  '#2c694e',
        'secondary-container':      '#dceee5',
        surface:                    '#f7f9fb',
        'surface-variant':          '#e4eaf0',
        'on-surface':               '#162033',
        'on-surface-variant':       '#657287',
        'surface-container-lowest': '#ffffff',
        'surface-container-low':    '#f1f5f8',
        'surface-container':        '#e8eef3',
        'surface-container-high':   '#d7e0e8',
        outline:                    '#8391a3',
        'outline-variant':          '#c8d2dc',
        error:                      '#b42318',
        'error-container':          '#fee4e2',
        'sidebar-bg':               '#041632',
        'sidebar-active':           '#12345f',
        'sidebar-active-indicator': '#2c694e',
      },
    },
  },
  plugins: [],
}
