import { Tile, Image } from 'views/HomePage/components/ImageTile/ImageTile.styles';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';

const ImageTile = ({ id, content }) => {
  return (
    <Tile>
      <Image>
        <Link to={{ pathname: `/ImageDetails/${id}` }}>
          <img src={content} alt="KymaImage" />
        </Link>
      </Image>
    </Tile>
  );
};

ImageTile.propTypes = {
  id: PropTypes.string.isRequired,
  content: PropTypes.string.isRequired,
};

export default ImageTile;
