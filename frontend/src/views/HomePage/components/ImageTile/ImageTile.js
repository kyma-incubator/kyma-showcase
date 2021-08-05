import { Tile, Image } from 'views/HomePage/components/ImageTile/ImageTile.styles';
import { Link } from 'react-router-dom';

const ImageTile = ({ url, id }) => {
  return (
    <Tile>
      <Image>
        <Link to={{ pathname: `/ImageDetails/${id}` }}>
          <img src={url} alt="KymaImage" />
        </Link>
        <p>{id}</p>
      </Image>
    </Tile>
  );
};

export default ImageTile;
