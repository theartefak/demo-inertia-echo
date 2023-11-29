<script setup>
import { Link } from '@inertiajs/vue3';
</script>

<template>
    <div class="flex min-h-screen flex-col items-center">
        <div class="flex w-full flex-grow items-center justify-center">
            <main
                class="my-16 w-full bg-white px-6 py-12 shadow-sm ring-1 ring-gray-950/5 dark:bg-gray-900 dark:ring-white/10 sm:max-w-lg sm:rounded-xl sm:px-12">
                <section class="grid auto-cols-fr gap-y-6">
                    <header>
                        <div class="mb-4 flex justify-center">
                            <div
                                class="text-xl font-bold leading-5 tracking-tight text-gray-950 dark:text-white">
                                <Link href="/">Artefak</Link>
                            </div>
                        </div>

                        <h1
                            class="text-center text-2xl font-bold tracking-tight text-gray-950 dark:text-white">
                            Sign in
                        </h1>
                    </header>

                    <slot />
                </section>
            </main>
        </div>
    </div>
</template>
