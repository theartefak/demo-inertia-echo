<script setup>
import { Head, Link } from '@inertiajs/vue3';
import Guest from '@/Layouts/Guest.vue';
defineOptions({ layout: Guest })
</script>

<template>
    <Head title="Welcome" />
    <h1>Welcome</h1>
    <Link :href="route('welcome')">Halo</Link>
</template>
