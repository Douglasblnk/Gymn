<script setup>
const props = defineProps({
  type: {
    type: String,
    default: 'button',
  },

  label: {
    type: String,
    default: '',
  },

  primary: Boolean,
  secondary: Boolean,
  ascendent: Boolean,
  disabled: Boolean,
  loading: Boolean,
  circle: Boolean,
  unelevated: Boolean,
})

const customClasses = computed(() => ({
  'g-button--primary': props.primary,
  'g-button--secondary': props.secondary,
  'g-button--ascendent': props.ascendent,
  'g-button--disabled': props.disabled,
  'g-button--loading': props.loading,
  'g-button--circle': props.circle,
  'g-button--unelevated': props.unelevated,
}))
</script>

<template>
  <button
    class="g-button"
    un-rounded-xl
    un-bg-gray-200
    un-py-3
    un-px-6
    un-transition-shadow
    un-shadow-2
    un-outline-none
    un-border-none
    un-active:shadow-8
    :class="customClasses"
    :disabled="disabled || loading"
    :type="type"
  >
    <slot>
      <span
        un-text="white"
        un-font="semibold tracking-wide"
        un-select="none"
      >
        {{ label }}
      </span>
    </slot>
  </button>
</template>

<style lang="sass">
.g-button
  &--primary
    background: var(--color-primary)

  &--secondary
    @apply bg-secondary

  &--ascendent
    @apply bg-ascendent

  &--disabled
    @apply opacity-30

  &--circle
    @apply rounded-full p-0

  &--unelevated
    @apply shadow-none
</style>
