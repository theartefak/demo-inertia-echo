<script setup>
import { Head, useForm } from '@inertiajs/vue3';
import Guest from '@/Layouts/Guest.vue';
import Input from '@/Components/Input.vue';
import Button from '@/Components/Button.vue';
import Checkbox from '@/Components/Checkbox.vue';

defineOptions({ layout: Guest });

const form = useForm({
    email: '',
    password: '',
    remember: false,
});
</script>

<template>
    <Head title="Welcome" />
    <div class="grid auto-cols-fr gap-y-6">
        <Input
            id="email"
            type="email"
            label="Email"
            v-model="form.email"
            autofocus
            autocomplete="username" />
        <Input
            id="username"
            type="text"
            label="Username"
            v-model="form.email"
            autocomplete="username" />
        <Checkbox
            id="remember"
            label="Remember me"
            name="remember"
            v-model:checked="form.remember" />
        <Button>Sign In</Button>
    </div>
</template>
