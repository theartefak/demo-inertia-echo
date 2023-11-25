import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';

export default defineConfig({
    plugins: [vue({
        template: {
            transformAssetUrls: {
                base: null,
                includeAbsolute: false,
            },
        },
    })],
    build: {
        manifest: 'manifest.json',
        outDir: resolve(__dirname, 'public'),
        rollupOptions: {
            input: 'resources/js/app.js',
        },
    },
    server: {
        cors: true,
    },
    resolve: {
        alias: {
            '@': resolve(__dirname, 'resources/js'),
        },
    },
});
