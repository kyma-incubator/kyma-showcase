import React, { useState } from 'react';
import { getAllImagesFromAPI } from 'API';
export const ImagesContext = React.createContext();

export function ImagesContextProvider({ children }) {
  const [images, setImages] = useState([]);

  const getImages = async () => setImages(await getAllImagesFromAPI());

  return <ImagesContext.Provider value={{ images, getImages }}>{children}</ImagesContext.Provider>;
}
