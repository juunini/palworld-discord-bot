export async function patchConfig(newConfig, oldConfig) {
  const changedConfig = Object.keys(oldConfig).reduce((acc, key) => {
    if (newConfig[key] !== oldConfig[key]) {
      acc[key] = newConfig[key]
    }

    return acc
  }, {});

  return fetch('/config', {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(changedConfig),
  })
    .catch(error => alert('Error:', error))
    .finally(() => location.reload());
}
