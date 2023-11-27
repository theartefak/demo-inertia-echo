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
        primary: 'bg-primary text-primary hover:bg-primary-500 focus:bg-primary-500 active:bg-primary-700',
        info: 'bg-info text-info hover:bg-info-500 focus:bg-info-500 active:bg-info-700',
        warning: 'bg-warning text-warning hover:bg-warning-500 focus:bg-warning-500 active:bg-warning-700',
        danger: 'bg-danger text-danger hover:bg-danger-500 focus:bg-danger-500 active:bg-danger-700',
        success: 'bg-success text-success hover:bg-success-500 focus:bg-success-500 active:bg-success-700',
    }[props.color];
});
</script>

<template>
    <button
        :type="props.type"
        :class="colorClass"
        class="inline-flex items-center px-4 py-2 border border-transparent rounded-md font-semibold text-xs uppercase tracking-widest focus:outline-none transition ease-in-out duration-150"
    >
        <slot />
    </button>
</template>
