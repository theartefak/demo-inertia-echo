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
        primary: 'bg-primary text-primary hover:bg-primary-600 focus:bg-primary-600 active:bg-primary-800',
        info: 'bg-info text-info hover:bg-info-600 focus:bg-info-600 active:bg-info-800',
        warning: 'bg-warning text-warning hover:bg-warning-600 focus:bg-warning-600 active:bg-warning-800',
        danger: 'bg-danger text-danger hover:bg-danger-600 focus:bg-danger-600 active:bg-danger-800',
        success: 'bg-success text-success hover:bg-success-600 focus:bg-success-600 active:bg-success-800',
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
