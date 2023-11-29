<template>
    <div style="--cols-default: repeat(1, minmax(0, 1fr))" class="grid gap-y-6">
        <div class="grid grid-cols-[--cols-default] gap-6">
            <slot />
        </div>
    </div>
</template>
