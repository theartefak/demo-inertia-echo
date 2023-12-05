<script setup>
import { Link } from '@inertiajs/vue3';
</script>

<template>
    <aside
        class="fixed inset-y-0 start-0 z-30 grid h-screen w-[--sidebar-width] -translate-x-full content-start bg-white transition-all rtl:translate-x-full dark:bg-gray-900 lg:sticky lg:z-0 lg:translate-x-0 lg:bg-transparent lg:shadow-none lg:ring-0 lg:transition-none rtl:lg:-translate-x-0 dark:lg:bg-transparent">
        <div class="overflow-x-clip">
            <header
                class="flex h-16 items-center bg-white px-6 ring-1 ring-gray-950/5 dark:bg-gray-900 dark:ring-white/10 lg:shadow-sm">
                <div>
                    <Link href="/">Artefak</Link>
                </div>
            </header>
        </div>

        <nav
            class="flex flex-col gap-y-7 overflow-y-auto overflow-x-hidden px-6 py-8">
            <ul class="flex flex-col gap-y-1">
                <li>
                    <Link
                        href="/"
                        class="relative flex items-center justify-center gap-x-3 rounded-lg px-2 py-2 text-sm outline-none transition duration-75 hover:bg-gray-100 focus-visible:bg-gray-100 dark:hover:bg-white/5 dark:focus-visible:bg-white/5"
                        >Dashboard</Link
                    >
                </li>
            </ul>
        </nav>
    </aside>
</template>
