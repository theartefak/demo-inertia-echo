<script setup>
import { Link } from '@inertiajs/vue3';
</script>

<template>
     <div class="min-h-screen flex flex-col sm:justify-center items-center pt-6 sm:pt-0 bg-gray-100 dark:bg-gray-950">
        <div>
            <Link href="/">
                Artefak
            </Link>
        </div>

        <div
            class="my-16 w-full bg-white px-6 py-12 shadow-sm ring-1 ring-gray-950/5 dark:bg-gray-900 dark:ring-white/10 sm:max-w-lg sm:rounded-xl sm:px-12"
        >
            <slot />
        </div>
    </div>
</template>
