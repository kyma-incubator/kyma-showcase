import { Loader } from 'assets/styles/style.js';
import PropTypes from 'prop-types';
import { ImageArea } from './ImageDetailsArea.styles.js';
import { DetailsContext } from 'contexts/detailsContext';

const ImageDetailsArea = ({ content }) => {
  return (
    <ImageArea>
      <img src={content} alt="Uploaded photo" />
    </ImageArea>
  );
};

ImageDetailsArea.propTypes = {
  id: PropTypes.string.isRequired,
};

export default ImageDetailsArea;
