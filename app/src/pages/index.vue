<script setup>
const email = ref()
</script>

<template>
  <div
    un-flex="~ col"
    un-justify-between
    un-h-screen
  >
    <div>
      <div
        un-flex
        un-justify-center
        un-items-center
        un-h-20rem
      >
        <GTitle />
      </div>

      <div
        un-flex="~ col"
        un-justify-center
        un-px-md
      >
        <GInputText
          v-model="email"
          label="Email"
          upper-label="Para começar, insira seu e-mail ou cadastre-se"
          icon="i-mdi-email"
          un-w-full
          clearable
        />

        <GButton
          label="Continuar"
          primary
          loading
          un-mt-md
        />
      </div>
    </div>

    <span
      un-text-gray-text
      un-flex
      un-w-full
      un-justify-center
      un-py-md
    >
      Não possui uma conta ainda?

      <span
        un-ml-xs
        un-text-secondary
      >
        Cadastre-se
      </span>
    </span>
  </div>
</template>
