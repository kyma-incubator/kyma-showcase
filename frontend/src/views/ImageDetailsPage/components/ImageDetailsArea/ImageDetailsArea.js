import { getImageDetailsFromAPI } from 'API.js';
import { Loader } from 'assets/styles/style.js';
import PropTypes from 'prop-types';
import { useEffect, useState } from 'react';
import { ImageArea } from './ImageDetailsArea.styles.js';

const ImageDetailsArea = ({ id }) => {
  const [imageDetails, setImageDetails] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    const callAPI = async () => {
      setIsLoading(true);
      try {
        setImageDetails(await getImageDetailsFromAPI(id));
        setIsLoading(false);
      } catch (err) {
        setIsLoading(false);
        setErrorMessage('Internal server error');
        console.error(err);
      }
    };
    callAPI();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <ImageArea>
      <p>{errorMessage}</p>
      {isLoading && <Loader />}
      {!isLoading && <img src={imageDetails.content} alt="Uploaded photo" />}
    </ImageArea>
  );
};

ImageDetailsArea.propTypes = {
  id: PropTypes.string.isRequired,
};

export default ImageDetailsArea;
