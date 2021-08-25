import { ImageArea, Img, H2 } from './ImageDetailsArea.styles.js';
import { useEffect, useState } from 'react';
import { getImageDetailsFromAPI } from 'API.js';

const ImageDetailsArea = ({ id }) => {
  const [imageDetails, setImageDetails] = useState({});
  useEffect(() => {
    const callAPI = async () => {
      try {
        setImageDetails(await getImageDetailsFromAPI(id));
      } catch (err) {
        console.error(err);
      }
    };
    callAPI();
  }, []);
  console.log(imageDetails);
  return (
    <ImageArea>
      <H2>Image title</H2>
      <Img src={imageDetails.content} alt="here will be" />
    </ImageArea>
  );
};

export default ImageDetailsArea;
