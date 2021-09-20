import PropTypes from 'prop-types';
import { ImageArea } from './ImageDetailsArea.styles.js';

const findNsfwTag = (gcp, key) => {
  return gcp?.find((obj) => Object.keys(obj).includes(key));
};

const ImageDetailsArea = ({ content, gcp }) => {
  gcp = gcp?.map(JSON.parse);
  const nsfw = findNsfwTag(gcp, 'nsfw')?.nsfw;
  return (
    <ImageArea nsfw={nsfw}>
      {nsfw && <p>Nsfw</p>}
      <img src={content} alt="Uploaded content" />
    </ImageArea>
  );
};

ImageDetailsArea.propTypes = {
  content: PropTypes.string.isRequired,
  gcp: PropTypes.array.isRequired,
};

export default ImageDetailsArea;
