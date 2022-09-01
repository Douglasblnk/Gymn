const errors = ref({})

const rules = {
  isRequired: {
    validator: value => !!value,
    msg: 'Este campo é obrigatório',
  },
  isValidEmail: {
    validator: value => /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/g.test(value),
    msg: 'E-mail inválido',
  },
  validPassword: {
    validator: (value) => {
      const strongRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/
      const mediumRegex = /^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})/

      if (value.length < 8) { return false }
      if (strongRegex.test(value)) { return 'green' }
      if (mediumRegex.test(value)) { return 'orange' }
      return 'red'
    },
  },
}

const useFieldModel = (models) => {
  errors.value = {}

  return models.map((model) => {
    errors.value[model] = ''

    return ref()
  })
}

export default (schema) => {
  const timeout = ref({})

  const setErrors = (key, msg) => {
    errors.value[key] = msg

    clearTimeout(timeout.value[key])

    timeout.value[key] = setTimeout(() => {
      errors.value[key] = ''
    }, 3000)
  }

  const validate = (models) => {
    return Object.entries(models).forEach(([ key, model ]) => {
      const validations = schema[key]

      const validationResult = validations.every((validation) => {
        const { validator, msg } = rules[validation]
        const result = validator(model.value)

        if (result === false) {
          setErrors(key, !msg && result !== Boolean ? result : msg)
        }

        return result
      })

      return validationResult
    })
  }

  return {
    useFieldModel,
    validate,
    errors,
  }
}
