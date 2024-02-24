export async function patchConfig(changedConfig) {
  return fetch('/config', {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(changedConfig),
  })
    .then(res => res.json())
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
