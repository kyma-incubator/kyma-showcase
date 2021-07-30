import { UploadedImage, Image } from 'views/HomePage/components/UploadedImage/UploadedImage.styles';
import { Link } from 'react-router-dom';

const UploadedImageBox = ({ url, cityName }) => {
  return (
    <UploadedImage>
      <Image>
        <Link to={{ pathname: `/ImageDetails/${cityName}` }}>
          <img src={url} alt="cat" />
        </Link>
        <p>{cityName}</p>
      </Image>
    </UploadedImage>
  );
};

export default UploadedImageBox;
