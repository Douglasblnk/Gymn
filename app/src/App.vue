<script setup>

</script>

<template>
  <RouterView v-slot="{ Component }">
    <Suspense>
      <transition
        name="fade"
        mode="out-in"
      >
        <component :is="Component" />
      </transition>
    </Suspense>
  </RouterView>
</template>
