export const flattenObj = (ob, separator = '.') => {
  const result = {}

  for (const i in ob) {
    if (typeof ob[i] === 'object' && !Array.isArray(ob[i])) {
      const temp = flattenObj(ob[i], separator)
      for (const j in temp) {
        result[`${ i }${ separator }${ j }`] = temp[j]
      }
    }

    else {
      result[i] = ob[i]
    }
  }

  return result
}
