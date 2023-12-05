<script setup>
import Sidebar from '@/Layouts/_partials/Sidebar.vue';
import Content from '@/Layouts/_partials/Content.vue';
</script>

<template>
    <div class="flex min-h-screen w-full overflow-x-clip">
        <Sidebar />
        <Content>
            <slot />
        </Content>
    </div>
</template>
