import { ImageArea, Img, H2 } from './ImageDetailsArea.styles.js';
import { useEffect, useState } from 'react';
import { getImageDetailsFromAPI } from 'API.js';

const ImageDetailsArea = ({ id }) => {
  const [imageDetails, setImageDetails] = useState({});
  const [isLoading, setIsLoading] = useState(true);
  useEffect(() => {
    const callAPI = async () => {
      setIsLoading(true);
      try {
        setImageDetails(await getImageDetailsFromAPI(id));
      } catch (err) {
        console.error(err);
        setIsLoading(false);
        //todo?? 
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

ImageDetailsArea.propTypes = {
  id: PropTypes.string.isRequired,
};

export default ImageDetailsArea;
