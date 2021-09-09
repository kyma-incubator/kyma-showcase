import React, { useState } from 'react';
import { getAllImagesFromAPI } from 'API';
export const ImagesContext = React.createContext();

export function ImagesContextProvider({ children }) {
  const [images, setImages] = useState([]);

  const getImages = async () => {
    let allImages = await getAllImagesFromAPI();
    allImages = allImages?.sort((d1, d2) => new Date(d1.time) - new Date(d2.time));
    setImages(allImages);
  };

  return <ImagesContext.Provider value={{ images, getImages }}>{children}</ImagesContext.Provider>;
}
