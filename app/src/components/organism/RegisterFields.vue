<script setup>
const { replace } = useRouter()

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

const register = () => {
  if (requestLoading.value) { return }

  const isValid = validate({ name, user, email, password })

  if (isValid) {
    return executeRequest({
      uri: 'register',
      data: {
        name: name.value,
        user: user.value,
        email: email.value,
        password: password.value,
      },
    })
  }
}
</script>

<template>
  <div
    un-flex="~ col"
    un-justify-center
    un-px-md
  >
    <GInputText
      v-model="name"
      label="Nome"
      name="name"
      icon="i-mdi-account"
      un-w-full
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

    <GInputPassword
      v-model="password"
      :error-msg="errors.password"
      un-w-full
      clearable
    />

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
</template>
