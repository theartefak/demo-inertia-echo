<script setup>
import { onMounted, ref } from 'vue';

defineProps({
    modelValue: {
        type: String,
        required: true,
    },
    label: {
        type: String,
        default: '',
    },
});

defineOptions({ inheritAttrs: false });

defineEmits(['update:modelValue']);

const input = ref(null);

onMounted(() => {
    if (input.value.hasAttribute('autofocus')) {
        input.value.focus();
    }
});

defineExpose({ focus: () => input.value.focus() });
</script>

<template>
    <div class="grid gap-y-2">
        <div v-if="label != ''" class="flex items-center justify-between gap-x-3">
            <label
                class="fi-fo-field-wrp-label inline-flex items-center gap-x-3"
                :for="$attrs.id"
            >
                <span class="text-sm font-medium leading-6 text-gray-950 dark:text-white">
                    {{ label }}
                    <sup
                        v-if="$attrs.required == '' || $attrs.required == 'true'"
                        class="text-danger-600 dark:text-danger-400 font-medium"
                    >*</sup>
                </span>
            </label>
        </div>

        <div class="grid gap-y-2">
            <div
                class="flex rounded-lg shadow-sm ring-1 transition duration-75 bg-white focus-within:ring-2 dark:bg-white/5 ring-gray-950/10 focus-within:ring-primary-600 dark:ring-white/20 dark:focus-within:ring-primary-500 fi-fo-text-input overflow-hidden"
            >
                <div class="min-w-0 flex-1">
                    <input
                        class="block w-full border-none py-1.5 text-base text-gray-950 transition duration-75 placeholder:text-gray-400 focus:ring-0 disabled:text-gray-500 disabled:[-webkit-text-fill-color:theme(colors.gray.500)] disabled:placeholder:[-webkit-text-fill-color:theme(colors.gray.400)] dark:text-white dark:placeholder:text-gray-500 dark:disabled:text-gray-400 dark:disabled:[-webkit-text-fill-color:theme(colors.gray.400)] dark:disabled:placeholder:[-webkit-text-fill-color:theme(colors.gray.500)] sm:text-sm sm:leading-6 bg-white/0 ps-3 pe-3"
                        :value="modelValue"
                        @input="$emit('update:modelValue', $event.target.value)"
                        v-bind="$attrs"
                        ref="input" />
                </div>
            </div>
        </div>
    </div>
</template>
