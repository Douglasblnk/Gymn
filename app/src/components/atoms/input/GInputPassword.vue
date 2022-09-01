<script setup>
const props = defineProps({
  modelValue: String,
})

const emit = defineEmits([ 'update:model-value' ])

const { validate } = useField({
  password: [ 'isRequired', 'validPassword' ],
})

const model = computed({
  get: () => props.modelValue,
  set: (value) => {
    emit('update:model-value', value)
  },
})

const STRONG_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/
const MEDIUM_REGEX = /^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})/

const isPasswordVisible = ref(false)

const setPasswordVisibility = () => {
  isPasswordVisible.value = !isPasswordVisible.value
}

const passwordStrength = computed(() => {
  if (STRONG_REGEX.test(model.value)) { return 'bg-green' }
  if (MEDIUM_REGEX.test(model.value)) { return 'bg-orange' }
  return 'bg-red'
})
</script>

<template>
  <GInputText
    v-model="model"
    label="Senha"
    name="password"
    :type="isPasswordVisible ? 'text' : 'password'"
  >
    <template #append>
      <Transition
        name="fade"
        mode="out-in"
      >
        <div
          v-if="model"
          un-w-5
          un-h-5
          un-rounded-full
          :class="passwordStrength"
        />
      </Transition>

      <GIcon
        :icon="isPasswordVisible ? 'i-mdi-eye' : 'i-mdi-lock'"
        un-text-white
        un-text-md
        un-cursor-pointer
        @click="setPasswordVisibility"
      />
    </template>
  </GInputText>
</template>
