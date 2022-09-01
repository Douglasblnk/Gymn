<script setup>
const makeLogin = () => {
  // todo
}
</script>

<template>
  <div
    un-flex="~ col"
    un-h-screen
  >
    <GTitle
      text="GYMN!"
      un-flex-grow
    />

    <SignInFields
      un-flex-grow
      @make-login="makeLogin"
    />

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
