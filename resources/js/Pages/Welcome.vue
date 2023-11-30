<script setup>
import { Head, useForm } from '@inertiajs/vue3';
import Guest from '@/Layouts/Guest.vue';
import Input from '@/Components/Input.vue';
import Button from '@/Components/Button.vue';
import Checkbox from '@/Components/Checkbox.vue';
import SimpleWrapForm from '@/Components/Wrapper/SimpleWrapForm.vue';
import SimpleWrapField from '@/Components/Wrapper/SimpleWrapField.vue';

defineOptions({ layout: Guest });

const form = useForm({
    email: '',
    password: '',
    remember: false,
});
</script>

<template>
    <Head title="Welcome" />
    <form method="post" class="grid auto-cols-fr gap-y-6">
        <SimpleWrapForm>
            <SimpleWrapField>
                <Input
                    id="email"
                    type="email"
                    label="Email"
                    v-model="form.email"
                    autofocus
                    autocomplete="username" />
            </SimpleWrapField>
            <SimpleWrapField>
                <Input
                    id="username"
                    type="text"
                    label="Username"
                    v-model="form.email"
                    autocomplete="username" />
            </SimpleWrapField>
            <SimpleWrapField>
                <Checkbox
                    id="remember"
                    label="Remember me"
                    name="remember"
                    v-model:checked="form.remember" />
            </SimpleWrapField>

            <Button>Sign In</Button>
        </SimpleWrapForm>
    </form>
</template>
