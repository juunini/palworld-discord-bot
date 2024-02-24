export async function patchConfig(changedConfig, successMessage, failedMessage) {
  return fetch('/config', {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(changedConfig),
  })
    .then(res => {
      console.log(res.status)
      if (res.status !== 202) {
        throw new Error(failedMessage);
      }

      alert(successMessage)
    })
    .catch((error) => alert(error))
    .finally(() => location.reload());
}

export function extractChangedConfig(newConfig, oldConfig) {
  return Object.keys(oldConfig).reduce((acc, key) => {
    if (newConfig[key] !== oldConfig[key]) {
      acc[key] = newConfig[key]
    }

    return acc
  }, {});
}
