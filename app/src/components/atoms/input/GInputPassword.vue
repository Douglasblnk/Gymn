<script setup>
const props = defineProps({
  modelValue: String,
})

const emit = defineEmits([ 'update:model-value' ])

const { useFieldModel, validate, errors } = useField({
  password: [ 'isRequired', 'validPassword' ],
})

const model = computed({
  get: () => props.modelValue,
  set: (value) => {
    emit('update:model-value', value)
  },
})
</script>

<template>
  <GInputText
    v-model="model"
    label="Senha"
    name="password"
    type="password"
    icon="i-mdi-lock"
  />
</template>
