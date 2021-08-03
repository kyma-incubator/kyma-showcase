export const APIGET = async () => {
  const response = await fetch('https://my-json-server.typicode.com/Lyczeq/images/images');

  return await response.json();
};

export const APIPOST = async (base) => {
  const response = await fetch('https://my-json-server.typicode.com/Lyczeq/images/images', {
    method: 'POST',
    body: JSON.stringify({
      base64: base,
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8',
    },
  });

  return  response.json();
};
