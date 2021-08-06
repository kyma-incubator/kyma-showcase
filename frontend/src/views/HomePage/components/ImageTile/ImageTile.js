import { Tile, Image } from 'views/HomePage/components/ImageTile/ImageTile.styles';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';

const ImageTile = ({ url, img }) => {
  return (
    <Tile>
      <Image>
        <Link to={{ pathname: `/ImageDetails/${img}` }}>
          <img src={url} alt="KymaImage" />
        </Link>
        <p>{img}</p>
      </Image>
    </Tile>
  );
};

ImageTile.propTypes = {
  url: PropTypes.string.isRequired,
  img: PropTypes.string.isRequired,
};

export default ImageTile;
