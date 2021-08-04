import { Tile, Image } from 'views/HomePage/components/ImageTile/ImageTile.styles';
import { Link } from 'react-router-dom';

const ImageTile = ({ url, cityName }) => {
  return (
    <Tile>
      <Image>
        <Link to={{ pathname: `/ImageDetails/${cityName}` }}>
          <img src={url} alt="cat" />
        </Link>
        <p>{cityName}</p>
      </Image>
    </Tile>
  );
};

export default ImageTile;
