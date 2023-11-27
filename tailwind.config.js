const colors = require('tailwindcss/colors');
import forms from '@tailwindcss/forms';

/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: 'class',
    content: ['./resources/**/*.{js,vue,html}'],
    theme: {
        extend: {
            textColor: {
                danger: {
                    DEFAULT: 'rgba(var(--danger-50), <alpha-value>)',
                },
                info: {
                    DEFAULT: 'rgba(var(--info-50), <alpha-value>)',
                },
                primary: {
                    DEFAULT: 'rgba(var(--primary-50), <alpha-value>)',
                },
                success: {
                    DEFAULT: 'rgba(var(--success-50), <alpha-value>)',
                },
                warning: {
                    DEFAULT: 'rgba(var(--warning-50), <alpha-value>)',
                },
            },
        },
        colors: {
            transparent: 'transparent',
            current: 'currentColor',
            black: colors.black,
            white: colors.white,
            danger: {
                DEFAULT: 'rgba(var(--danger-600), <alpha-value>)',
                50: 'rgba(var(--danger-50), <alpha-value>)',
                100: 'rgba(var(--danger-100), <alpha-value>)',
                200: 'rgba(var(--danger-200), <alpha-value>)',
                300: 'rgba(var(--danger-300), <alpha-value>)',
                400: 'rgba(var(--danger-400), <alpha-value>)',
                500: 'rgba(var(--danger-500), <alpha-value>)',
                600: 'rgba(var(--danger-600), <alpha-value>)',
                700: 'rgba(var(--danger-700), <alpha-value>)',
                800: 'rgba(var(--danger-800), <alpha-value>)',
                900: 'rgba(var(--danger-900), <alpha-value>)',
                950: 'rgba(var(--danger-950), <alpha-value>)',
            },
            gray: {
                50: 'rgba(var(--gray-50), <alpha-value>)',
                100: 'rgba(var(--gray-100), <alpha-value>)',
                200: 'rgba(var(--gray-200), <alpha-value>)',
                300: 'rgba(var(--gray-300), <alpha-value>)',
                400: 'rgba(var(--gray-400), <alpha-value>)',
                500: 'rgba(var(--gray-500), <alpha-value>)',
                600: 'rgba(var(--gray-600), <alpha-value>)',
                700: 'rgba(var(--gray-700), <alpha-value>)',
                800: 'rgba(var(--gray-800), <alpha-value>)',
                900: 'rgba(var(--gray-900), <alpha-value>)',
                950: 'rgba(var(--gray-950), <alpha-value>)',
            },
            info: {
                DEFAULT: 'rgba(var(--info-600), <alpha-value>)',
                50: 'rgba(var(--info-50), <alpha-value>)',
                100: 'rgba(var(--info-100), <alpha-value>)',
                200: 'rgba(var(--info-200), <alpha-value>)',
                300: 'rgba(var(--info-300), <alpha-value>)',
                400: 'rgba(var(--info-400), <alpha-value>)',
                500: 'rgba(var(--info-500), <alpha-value>)',
                600: 'rgba(var(--info-600), <alpha-value>)',
                700: 'rgba(var(--info-700), <alpha-value>)',
                800: 'rgba(var(--info-800), <alpha-value>)',
                900: 'rgba(var(--info-900), <alpha-value>)',
                950: 'rgba(var(--info-950), <alpha-value>)',
            },
            primary: {
                DEFAULT: 'rgba(var(--primary-600), <alpha-value>)',
                50: 'rgba(var(--primary-50), <alpha-value>)',
                100: 'rgba(var(--primary-100), <alpha-value>)',
                200: 'rgba(var(--primary-200), <alpha-value>)',
                300: 'rgba(var(--primary-300), <alpha-value>)',
                400: 'rgba(var(--primary-400), <alpha-value>)',
                500: 'rgba(var(--primary-500), <alpha-value>)',
                600: 'rgba(var(--primary-600), <alpha-value>)',
                700: 'rgba(var(--primary-700), <alpha-value>)',
                800: 'rgba(var(--primary-800), <alpha-value>)',
                900: 'rgba(var(--primary-900), <alpha-value>)',
                950: 'rgba(var(--primary-950), <alpha-value>)',
            },
            success: {
                DEFAULT: 'rgba(var(--success-600), <alpha-value>)',
                50: 'rgba(var(--success-50), <alpha-value>)',
                100: 'rgba(var(--success-100), <alpha-value>)',
                200: 'rgba(var(--success-200), <alpha-value>)',
                300: 'rgba(var(--success-300), <alpha-value>)',
                400: 'rgba(var(--success-400), <alpha-value>)',
                500: 'rgba(var(--success-500), <alpha-value>)',
                600: 'rgba(var(--success-600), <alpha-value>)',
                700: 'rgba(var(--success-700), <alpha-value>)',
                800: 'rgba(var(--success-800), <alpha-value>)',
                900: 'rgba(var(--success-900), <alpha-value>)',
                950: 'rgba(var(--success-950), <alpha-value>)',
            },
            warning: {
                DEFAULT: 'rgba(var(--warning-600), <alpha-value>)',
                50: 'rgba(var(--warning-50), <alpha-value>)',
                100: 'rgba(var(--warning-100), <alpha-value>)',
                200: 'rgba(var(--warning-200), <alpha-value>)',
                300: 'rgba(var(--warning-300), <alpha-value>)',
                400: 'rgba(var(--warning-400), <alpha-value>)',
                500: 'rgba(var(--warning-500), <alpha-value>)',
                600: 'rgba(var(--warning-600), <alpha-value>)',
                700: 'rgba(var(--warning-700), <alpha-value>)',
                800: 'rgba(var(--warning-800), <alpha-value>)',
                900: 'rgba(var(--warning-900), <alpha-value>)',
                950: 'rgba(var(--warning-950), <alpha-value>)',
            },
        },
    },
    plugins: [forms],
};
