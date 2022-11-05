<script setup>
const {
  executeRequest,
  clearRequest,
  requestResult,
  requestLoading,
} = useRequest({ showAlertOnError: true })

const { useFieldModel, validate, errors } = useField({
  email: [ 'isRequired', 'isValidEmail' ],
  password: [ 'isRequired' ],
})

const [ email, password ] = useFieldModel([ 'email', 'password' ])

const isEmailValid = computed(() => {
  return requestResult.value === true
})

const upperLabel = computed(() => {
  return isEmailValid.value
    ? 'Insira sua senha'
    : 'Para começar, insira seu e-mail ou cadastre-se'
})

const checkEmail = async () => {
  if (requestLoading.value) { return }

  validate({ email }) && executeRequest({
    uri: 'validate-email',
    data: {
      email: email.value,
    },
  })
}

const makeLogin = async () => {
  validate({ password }) && executeRequest({
    uri: 'sign-in',
    data: {
      email: email.value,
      password: password.value,
    },
  })
}

const executeAction = () => {
  isEmailValid.value ? makeLogin() : checkEmail()
}

const testCookie = () => {
  executeRequest({
    uri: 'student',
    method: 'GET',
  })
}
</script>

<template>
  <div
    un-w="sm:col-7 md:col-5 lg:col-3"
    un-m-auto
    un-flex="~ col"
    un-h-screen
    un-overflow-hidden
  >
    <GTitle
      text="GYMN!"
      un-flex-grow
    />

    <div
      un-flex="~ col"
      un-flex-grow
      un-justify-center
      un-px-md
    >
      <div
        un-mb-4xl
        un-text-center
      >
        <Transition
          name="fade"
          mode="out-in"
        >
          <span
            v-if="!isEmailValid"
            un-text="white 2xl"
          >
            Acesse sua conta
          </span>

          <div
            v-else
            un-column
            un-items-center
            un-space-y-md
          >
            <span un-text="white 2xl">
              Olá
            </span>

            <GChip
              v-if="isEmailValid"
              color="var(--color-secondary)"
              :text="email"
              clearable
              size="xl"
              @clear="clearRequest"
            />
          </div>
        </Transition>
      </div>

      <Transition
        name="slide-left"
        mode="out-in"
      >
        <GInputText
          v-if="!isEmailValid"
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

        <GInputText
          v-else
          v-model="password"
          label="Senha"
          name="password"
          type="password"
          icon="i-mdi-lock"
          :upper-label="upperLabel"
          :error-msg="errors.password"
          un-w-full
          clearable
          @keypress.enter="makeLogin"
        />
      </Transition>

      <GButton
        :label="isEmailValid ? 'Acessar' : 'Continuar'"
        :loading="requestLoading"
        un-mt-sm
        primary
        @click="executeAction"
      />

      <GButton
        label="cookie"
        @click="testCookie"
      />
    </div>

    <span
      un-flex-grow
      un-text-gray-text
      un-flex
      un-w-full
      un-justify-center
      un-items-end
      un-py-md
    >
      Não possui uma conta ainda?

      <span
        un-ml-xs
        un-text-secondary
        un-cursor-pointer
      >
        <RouterLink to="/registrar">
          Cadastre-se
        </RouterLink>
      </span>
    </span>
  </div>
</template>

<route lang="yaml">
alias: /login
meta:
  transition: slide-down
</route>
