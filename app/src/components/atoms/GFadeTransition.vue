<template>
  <Transition mode="out-in">
    <slot />
  </Transition>
</template>

<style lang="sass">
.v-enter-active,
.v-leave-active
  transition: opacity 0.5s ease

.v-enter-from,
.v-leave-to
  opacity: 0
</style>
