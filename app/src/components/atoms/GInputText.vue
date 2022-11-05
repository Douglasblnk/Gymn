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

  name: String,
  placeholder: String,
  upperLabel: String,
  label: String,
  icon: String,
  readonly: Boolean,
  clearable: Boolean,
  errorMsg: String,
  hideBottomSpace: Boolean,
})

const emit = defineEmits([ 'update:modelValue', 'clear' ])

const inputRef = ref()

const clearField = () => {
  emit('update:modelValue', '')
  emit('clear')
}

const model = computed({
  get: () => props.modelValue,
  set: value => emit('update:modelValue', value),
})
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
      un-z-1
      un-transition="duration-500 colors"
      :un-border="!!errorMsg ? '1 solid negative' : '1 solid transparent'"
      @click="inputRef.focus()"
    >
      <input
        v-model="model"
        v-bind="$attrs"
        ref="inputRef"
        un-p="t-6 b-3 x-6"
        un-w-full
        un-outline-none
        un-text-white
        un-font-bold
        un-border-none
        un-bg-transparent
        :type="type"
        :name="name"
        :placeholder="placeholder"
        :readonly="readonly"
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
        un-items-center
        un-space-x-sm
        un-px-md
      >
        <Transition
          name="fade"
          mode="out-in"
        >
          <GIcon
            v-if="clearable && model"
            un-text-white
            un-cursor-pointer
            un-text-2xl
            icon="i-mdi-close"
            @click.stop="clearField"
          />
        </Transition>

        <slot name="append" />

        <GIcon
          v-if="icon"
          un-text-white
          un-text-2xl
          :icon="icon"
        />
      </div>
    </div>

    <div
      v-if="!hideBottomSpace"
      un-ml-lg
      un-h-5
      un-flex
      un-items-center
    >
      <Transition
        name="slide-down"
        mode="out-in"
      >
        <div
          v-if="!!errorMsg"
          un-z="-1"
          un-text="negative sm middle"
        >
          {{ errorMsg }}
        </div>
      </Transition>
    </div>
  </div>
</template>

<style lang="sass">
.g-input-text
  label
    transform: translateY(-55%)
    transform-origin: left top
    user-select: none
    transition: transform .3s cubic-bezier(0.4, 0, 0.2, 1), color .3s  cubic-bezier(0.4, 0, 0.2, 1), top .3s

  input:focus + label
    transition: transform .3s
    color: var(--color-secondary)
    transform: translateY(-100%) scale(0.75)

  &--active
    transform: translateY(-100%) scale(0.75) !important
</style>
