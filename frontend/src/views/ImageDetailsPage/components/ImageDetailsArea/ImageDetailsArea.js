import PropTypes from 'prop-types';
import { ImageArea } from './ImageDetailsArea.styles.js';

const findObject = (gcp, key) => {
  return gcp?.find((obj) => Object.keys(obj).includes(key));
};
const ImageDetailsArea = ({ content, gcp }) => {
  gcp = gcp?.map(JSON.parse);
  console.log(gcp);
  const nsfw = findObject(gcp, 'nsfw')?.nsfw;
  console.log(nsfw);
  return (
    <ImageArea nsfw={nsfw}>
      {nsfw && <p>Nsfw</p>}
      <img src={content} alt="Uploaded content" />
    </ImageArea>
  );
};

ImageDetailsArea.propTypes = {
  content: PropTypes.string.isRequired,
};

export default ImageDetailsArea;
