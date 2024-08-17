import type { Config } from 'tailwindcss';

const config: Config = {
    content: [
        './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
        './src/components/**/*.{js,ts,jsx,tsx,mdx}',
        './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    ],
    theme: {
        extend: {
            primary: {
                light: '#A86854',
                DEFAULT: '#884A39',
            },
            secondary: {
                light: '#D29D73',
                DEFAULT: '#C38154',
            },
            tertiary: {
                light: '#FFD7A1',
                DEFAULT: '#FFC26F',
            },
            neutral: {
                light: '#FBF1D6',
                DEFAULT: '#F9E0BB',
            },
            destructive: {
                light: '#8B4C52',
                DEFAULT: '#6D2932',
            },
        },
    },
    plugins: [],
};
export default config;
