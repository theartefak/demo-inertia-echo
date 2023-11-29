<script setup>
import { computed } from 'vue';

const emit = defineEmits(['update:checked']);

defineOptions({ inheritAttrs: false });

const props = defineProps({
    checked: {
        type: [Array, Boolean],
        required: true,
    },
    value: {
        default: null,
    },
    label: {
        type: String,
        required: true,
    },
});

const proxyChecked = computed({
    get() {
        return props.checked;
    },

    set(val) {
        emit('update:checked', val);
    },
});
</script>

<template>
    <div class="flex items-center justify-between gap-x-3">
        <label class="inline-flex items-center gap-x-3" :for="$attrs.id">
            <input
                type="checkbox"
                :value="value"
                v-model="proxyChecked"
                class="rounded border-none bg-white text-primary-600 shadow-sm ring-1 ring-gray-950/10 transition duration-75 checked:ring-0 focus:ring-2 focus:ring-primary focus:ring-offset-0 checked:focus:ring-primary-500/50 disabled:pointer-events-none disabled:bg-gray-50 disabled:text-gray-50 disabled:checked:bg-current disabled:checked:text-gray-400 dark:bg-white/5 dark:text-primary-500 dark:ring-white/20 dark:checked:bg-primary-500 dark:focus:ring-primary-500 dark:checked:focus:ring-primary-400/50 dark:disabled:bg-transparent dark:disabled:ring-white/10 dark:disabled:checked:bg-gray-600"
                v-bind="$attrs" />

            <span
                class="cursor-pointer select-none text-sm font-medium leading-6 text-gray-950 dark:text-white">
                {{ label }}
            </span>
        </label>
    </div>
</template>
