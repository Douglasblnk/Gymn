<script setup>
const emit = defineEmits([ 'make-login' ])

const {
  executeRequest,
  clearRequest,
  requestResult,
  requestLoading,
  requestError,
} = useRequest()

const { watchError } = useWatch()
watchError(requestError)

const { useFieldModel, validate, errors } = useField({
  email: [ 'isRequired', 'isValidEmail' ],
  password: [ 'isRequired' ],
})

const [ email, password ] = useFieldModel([ 'email', 'password' ])

const checkEmail = () => {
  if (requestLoading.value) { return }

  const isValid = validate({ email })

  if (isValid) {
    return executeRequest({
      uri: 'email-validation',
      data: {
        email: email.value,
      },
    })
  }
}

const isEmailValid = computed(() => {
  return requestResult.value === true
})

const upperLabel = computed(() => {
  return isEmailValid.value
    ? 'Insira sua senha'
    : 'Para começar, insira seu e-mail ou cadastre-se'
})
</script>

<template>
  <div
    un-flex="~ col"
    un-justify-center
    un-px-md
  >
    <GInputText
      v-model="email"
      label="Email"
      name="email"
      icon="i-mdi-email"
      un-w-full
      :upper-label="upperLabel"
      :error-msg="errors.email"
      :readonly="isEmailValid"
      clearable
      @keypress.enter="checkEmail"
      @clear="clearRequest"
    />

    <Transition
      name="slide-down"
      mode="out-in"
    >
      <GInputText
        v-if="isEmailValid"
        v-model="password"
        key="password"
        label="Senha"
        name="password"
        type="password"
        icon="i-mdi-lock"
        :error-msg="errors.password"
        un-w-full
        clearable
        @keypress.enter="makeLogin"
      />
    </Transition>

    <GButton
      label="Continuar"
      :loading="requestLoading"
      un-mt-sm
      primary
      @click="checkEmail"
    />
  </div>
</template>
