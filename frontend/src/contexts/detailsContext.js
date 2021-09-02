import React, { useState } from 'react';
import { getImageDetailsFromAPI } from 'API';
export const DetailsContext = React.createContext();

export function DetailsContextProvider({ children }) {
  const [imageDetails, setImageDetails] = useState(null)
  const getImageDetails = async (id) => setImageDetails(await getImageDetailsFromAPI(id))

  return <DetailsContext.Provider value={{ imageDetails, getImageDetails }}>{children}</DetailsContext.Provider>;
}
