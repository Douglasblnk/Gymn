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
    validator: value => value.length >= 8,
    msg: 'Senha muito curta',
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

  const getRefValue = (model) => {
    return model?.__v_isRef && typeof model === 'object'
      ? model.value
      : model
  }

  const validate = (models) => {
    return Object.entries(models).map(([ key, model ]) => {
      const validations = schema[key]

      const validationResult = validations.every((validation) => {
        const { validator, msg } = rules[validation]
        const result = validator(getRefValue(model))

        if (result === false) {
          setErrors(key, msg)
        }

        return result
      })

      return validationResult
    }).every(hasError => hasError)
  }

  return {
    useFieldModel,
    validate,
    errors,
  }
}
