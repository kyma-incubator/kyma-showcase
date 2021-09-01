import { Loader } from 'assets/styles/style.js';
import PropTypes from 'prop-types';
import { useEffect, useState, useContext } from 'react';
import { ImageArea } from './ImageDetailsArea.styles.js';
import { DetailsContext } from 'contexts/detailsContext';

const ImageDetailsArea = ({ id }) => {
  const [isLoading, setIsLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');
  const { imageDetails, getImageDetails } = useContext(DetailsContext);

  useEffect(() => {
    const callAPI = async () => {
      setIsLoading(true);
      try {
        await getImageDetails(id);
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
