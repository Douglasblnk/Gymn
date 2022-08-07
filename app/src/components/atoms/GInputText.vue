<script setup>
const props = defineProps({
  modelValue: {
    type: [ String, Number ],
    default: '',
  },

  type: {
    type: String,
    default: 'text',
  },

  labelColor: {
    type: String,
    default: 'white',
  },

  placeholder: String,
  upperLabel: String,
  label: String,
  errorMsg: String,
  icon: String,
  error: Boolean,
  clearable: Boolean,
})

const emit = defineEmits([ 'update:modelValue' ])

const model = computed({
  get: () => props.modelValue,
  set: value => emit('update:modelValue', value),
})

const inputRef = ref()
</script>

<template>
  <div class="g-input-text">
    <p
      v-if="upperLabel"
      un-m-b-xs
      un-p-l-6
      un-text-14px
      :style="{ color: labelColor }"
    >
      {{ upperLabel }}
    </p>

    <div
      un-min-w-10
      un-relative
      un-flex
      un-justify-between
      un-items-center
      un-rounded-2xl
      un-bg-gray-bg
      un-cursor-text
      :un-border="error ? '1 solid negative' : 'none'"
      @click="inputRef.focus()"
    >
      <input
        ref="inputRef"
        v-model="model"
        un-p="t-7 b-3 x-6"
        un-w-full
        un-outline-none
        un-text-white
        un-font-bold
        un-border-none
        un-bg-transparent
        :type="type"
        :placeholder="placeholder"
      >

      <label
        un-absolute
        un-top-7
        un-left-6
        un-text-gray-text
        un-pointer-events-none
        :class="model ? 'g-input-text--active' : ''"
      >
        {{ label }}
      </label>

      <div
        un-flex
        un-space-x-sm
        un-px-md
      >
        <GFadeTransition>
          <GIcon
            v-if="clearable && model"
            un-text-white
            un-cursor-pointer
            un-text-md
            icon="i-mdi-close"
            @click.stop="emit('update:modelValue', '')"
          />
        </GFadeTransition>

        <slot name="append" />

        <GIcon
          v-if="icon"
          un-text-white
          un-text-md
          :icon="icon"
        />
      </div>
    </div>
  </div>
</template>

<style lang="sass">
.g-input-text
  label
    transform: translateY(-50%)
    transform-origin: left top
    user-select: none
    transition: transform .3s cubic-bezier(0.4, 0, 0.2, 1), color .3s  cubic-bezier(0.4, 0, 0.2, 1), top .3s

  input:focus + label
    transition: transform .3s
    color: var(--color-secondary)
    transform: translateY(-90%) scale(0.75)

  &--active
    transform: translateY(-90%) scale(0.75) !important
</style>
