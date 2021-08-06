export const APIGET = async (url) => {
  const response = await fetch(url);

  return response.json();
};

export const APIPOST = async (base, url, random) => {
  const response = await fetch(url, {
    method: 'POST',
    body: JSON.stringify({
      url: base,
      img: random.toString(),
    }),
    headers: {
      'Content-type': 'application/json',
    },
  });

  return response.json();
};
