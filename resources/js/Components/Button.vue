<script setup>
import { computed } from 'vue';

const props = defineProps({
    type: {
        type: String,
        default: 'button',
    },
    color: {
        type: String,
        default: 'primary',
    },
});

const colorClass = computed(() => {
    return {
        primary:
            'bg-primary text-primary hover:bg-primary-500 focus:bg-primary-500 active:bg-primary-700',
        info: 'bg-info text-info hover:bg-info-500 focus:bg-info-500 active:bg-info-700',
        warning:
            'bg-warning text-warning hover:bg-warning-500 focus:bg-warning-500 active:bg-warning-700',
        danger: 'bg-danger text-danger hover:bg-danger-500 focus:bg-danger-500 active:bg-danger-700',
        success:
            'bg-success text-success hover:bg-success-500 focus:bg-success-500 active:bg-success-700',
    }[props.color];
});
</script>

<template>
    <div>
        <div class="grid grid-cols-[repeat(auto-fit,minmax(0,1fr))] gap-3">
            <button
                :type="props.type"
                :class="colorClass"
                class="relative inline-grid grid-flow-col items-center justify-center gap-1.5 rounded-lg px-3 py-2 text-sm font-semibold shadow-sm outline-none transition duration-75 focus-visible:ring-2">
                <slot />
            </button>
        </div>
    </div>
</template>
