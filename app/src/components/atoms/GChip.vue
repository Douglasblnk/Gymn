<script setup>
defineProps({
  text: String,
  icon: String,
  clearable: Boolean,

  size: {
    type: String,
    default: 'sm',
  },

  color: {
    type: String,
    default: 'var(--color-primary)',
  },
})

const emit = defineEmits([ 'emit' ])
</script>

<template>
  <div
    un-flex
    un-w-fit
    un-justify-center
    un-items-center
    un-rounded-full
    un-text-white
    un-px-md
    un-py-sm
    :class="`text-${size}`"
    :style="`background: ${color}`"
  >
    <span>
      {{ text }}
    </span>

    <GIcon
      icon="i-mdi-close-circle"
      un-text-xl
      un-ml-sm
      un-opacity="70 hover:100"
      un-cursor-pointer
      @click="emit('clear')"
    />
  </div>
</template>
