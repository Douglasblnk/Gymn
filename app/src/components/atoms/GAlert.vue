<script setup>
defineProps({
  isShowing: Boolean,

  text: {
    type: String,
    default: '',
  },

  type: {
    type: String,
    default: 'warning',
    validator: type => [
      'warning', 'info', 'positive', 'negative',
    ].includes(type),
  },
})
</script>

<template>
  <Teleport to="body">
    <Transition
      name="slide-down"
      mode="in-out"
    >
      <div
        v-if="isShowing"
        un-pos="fixed top-lg"
        un-w-full
        un-flex
        un-justify-center
        un-z-notify
      >
        <div
          :class="`bg-${type}`"
          un-text-white
          un-shadow="10 gray-700"
          un-flex
          un-justify-center
          un-items-center
          un-px-lg
          un-py-sm
          un-rounded-full
          un-font-base
        >
          {{ text }}
        </div>
      </div>
    </Transition>
  </Teleport>
</template>
