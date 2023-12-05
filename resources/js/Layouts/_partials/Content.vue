<template>
    <div class="flex w-screen flex-1 flex-col">
        <div class="sticky top-0 z-20 overflow-x-clip">
            <nav
                class="flex h-16 items-center gap-x-4 bg-white px-4 shadow-sm ring-1 ring-gray-950/5 dark:bg-gray-900 dark:ring-white/10 md:px-6 lg:px-8">
                <button
                    class="relative -m-1.5 flex h-9 w-9 items-center justify-center rounded-lg text-gray-400 outline-none transition duration-75 hover:text-gray-500 focus-visible:ring-2 focus-visible:ring-primary-600 disabled:pointer-events-none disabled:opacity-70 dark:text-gray-500 dark:hover:text-gray-400 dark:focus-visible:ring-primary-500 lg:hidden">
                    <span class="sr-only">Expand sidebar</span>
                    <svg
                        class="h-6 w-6"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke-width="1.5"
                        stroke="currentColor"
                        aria-hidden="true">
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"></path>
                    </svg>
                </button>
                <div class="ms-auto flex items-center gap-x-4">
                    <div>Halo</div>
                </div>
            </nav>
        </div>
        <main class="mx-auto h-full w-full max-w-7xl px-4 md:px-6 lg:px-8">
            <slot />
        </main>
    </div>
</template>
