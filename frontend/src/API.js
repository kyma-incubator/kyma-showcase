export const APIGET = async (url) => {
  const response = await fetch(url);

  return response.json();
};

export const APIPOST = async (base, url) => {
  const response = await fetch(url, {
    method: 'POST',
    body: JSON.stringify({
      base64: base,
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8',
    },
  });

  return response.json();
};
