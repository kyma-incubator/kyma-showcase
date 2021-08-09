const API_URL = 'http://localhost:8081/v1/images';

export const getAllImagesFromAPI = async () => {
  const response = await fetch(API_URL);
  return response.json();
};

export const postImageToAPI = async (base64) => {
  const response = await fetch(API_URL, {
    method: 'POST',
    body: JSON.stringify({
      content: base64,
    }),
    headers: {
      'Content-type': 'application/json',
    },
  });

  return response.json();
};