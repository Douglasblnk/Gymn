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

  to: String,
  icon: String,
  primary: Boolean,
  secondary: Boolean,
  ascendent: Boolean,
  loading: Boolean,
  iconLeft: Boolean,
})

const customClasses = computed(() => ({
  's-button--primary': props.primary,
  's-button--secondary': props.secondary,
  's-button--ascendent': props.ascendent,
  'cursor-wait': props.loading,
}))
</script>

<template>
  <button
    class="s-button"
    un-rounded-2xl
    un-border-none
    un-text-white
    un-font-bold
    un-px-md
    un-py-3
    un-outline-none
    un-active:shadow-8
    un-flex
    un-transition
    un-items-center
    un-justify-center
    :class="customClasses"
    :type="type"
  >
    <GIcon
      v-if="loading"
      icon="i-mdi-loading"
      un-animate-spin
      un-text="xl"
    />

    <template v-else>
      <GIcon
        v-if="icon && iconLeft"
        :icon="icon"
        un-text="xl"
        un-mr-sm
      />
      <slot>
        <span
          un-font-tracking-wide
          un-select="none"
        >
          {{ label }}
        </span>
      </slot>

      <GIcon
        v-if="icon && !iconLeft"
        :icon="icon"
        un-text="xl"
        un-ml-sm
      />
    </template>
  </button>
</template>

<style lang="sass">
.s-button
  &--primary
    background: var(--color-primary)

  &--secondary
    background: var(--color-secondary)

  &--ascendent
    background: var(--color-ascendent)
</style>
