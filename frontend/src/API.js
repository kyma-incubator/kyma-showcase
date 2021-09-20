let API_CONFIG;

export const fetchAPIConfig = async () => {
  const response = await fetch('/config/config.json');
  API_CONFIG = await response.json();
};

export const fetchMapAPIConfig = async () => {
  return API_CONFIG.API_KEY_PARAM;
};

export const getAllImagesFromAPI = async () => {
  const response = await fetch(API_CONFIG.API_URL);
  return response.json();
};

export const postImageToAPI = async (base64) => {
  const response = await fetch(API_CONFIG.API_URL, {
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

export const getImageDetailsFromAPI = async (id) => {
  const response = await fetch(API_CONFIG.API_URL + `/${id}`);
  return response.json();
};
