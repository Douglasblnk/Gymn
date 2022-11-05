<script setup>
const { replace } = useRouter()
const { setAlert } = useAlert()

const {
  executeRequest,
  requestLoading,
} = useRequest({ showAlertOnError: true })

const { useFieldModel, validate, errors } = useField({
  name: [ 'isRequired' ],
  user: [ 'isRequired' ],
  email: [ 'isRequired', 'isValidEmail' ],
  password: [ 'isRequired', 'validPassword' ],
})

const [
  name,
  user,
  email,
  password,
] = useFieldModel([ 'name', 'user', 'email', 'password' ])

const STRONG_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/
const MEDIUM_REGEX = /^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})/

const isPasswordVisible = ref(false)

const passwordStrength = computed(() => {
  if (STRONG_REGEX.test(password.value)) { return 'bg-green' }
  if (MEDIUM_REGEX.test(password.value)) { return 'bg-orange' }
  return 'bg-red'
})

const setPasswordVisibility = () => {
  isPasswordVisible.value = !isPasswordVisible.value
}

const clearFields = () => {
  name.value = ''
  user.value = ''
  email.value = ''
  password.value = ''
}

const register = () => {
  if (requestLoading.value) { return }

  const isValid = validate({ name, user, email, password })

  if (isValid) {
    const success = executeRequest({
      uri: 'register',
      data: {
        name: name.value,
        user: user.value,
        email: email.value,
        password: password.value,
      },
    })

    success && setAlert({
      text: 'Cadastro realizado com sucesso',
      type: 'positive',
    })

    clearFields()
    replace('/login')
  }
}
</script>

<template>
  <div
    un-w="sm:col-7 md:col-5 lg:col-3"
    un-m-auto
    un-flex="~ col"
    un-h-screen
  >
    <GTitle
      text="Cadastrar."
      small
      un-flex-grow
    />

    <div
      un-flex="~ col"
      un-flex-grow
      un-justify-center
      un-px-md
    >
      <GInputText
        v-model="name"
        label="Nome"
        name="name"
        icon="i-mdi-account"
        un-w-full
        upper-label="Preencha os campos para se cadastrar"
        :error-msg="errors.name"
        clearable
      />

      <GInputText
        v-model="user"
        label="Usuário"
        name="user"
        icon="i-mdi-card-account-details"
        :error-msg="errors.user"
        un-w-full
        clearable
      />

      <GInputText
        v-model="email"
        label="E-mail"
        name="email"
        icon="i-mdi-email"
        :error-msg="errors.email"
        un-w-full
        clearable
      />

      <GInputText
        v-model="password"
        label="Senha"
        name="password"
        un-w-full
        clearable
        :type="isPasswordVisible ? 'text' : 'password'"
        :error-msg="errors.password"
        @keypress.enter="register"
      >
        <template #append>
          <Transition
            name="fade"
            mode="out-in"
          >
            <div
              v-if="password"
              un-w-5
              un-h-5
              un-rounded-full
              :class="passwordStrength"
            />
          </Transition>

          <GIcon
            :icon="isPasswordVisible ? 'i-mdi-eye' : 'i-mdi-lock'"
            un-text-white
            un-text-2xl
            un-cursor-pointer
            @click="setPasswordVisibility"
          />
        </template>
      </GInputText>

      <GButton
        label="Cadastrar"
        un-mt="2xl"
        :loading="requestLoading"
        primary
        @click="register"
      />

      <GButton
        label="Voltar"
        un-mt-sm
        :loading="requestLoading"
        un-bg-gray-bg
        @click="replace('/login')"
      />
    </div>
  </div>
</template>

<route lang="yaml">
alias: /registrar
meta:
  transition: slide-down
</route>
