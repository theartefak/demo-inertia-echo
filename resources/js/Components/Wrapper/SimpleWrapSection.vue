<template>
    <section class="flex flex-col gap-y-8 py-8">
        <div>
            <div class="grid flex-1 auto-cols-fr gap-y-8">
                <div class="flex flex-col gap-y-6">
                    <div>
                        <div
                            class="divide-y divide-gray-200 overflow-hidden rounded-xl bg-white p-5 shadow-sm ring-1 ring-gray-950/5 dark:divide-white/10 dark:bg-gray-900 dark:ring-white/10">
                            <slot />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
