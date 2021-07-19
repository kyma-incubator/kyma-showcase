import {
  UploadedImage,
  Image,
} from 'components/UploadedImage/UploadedImage.styles';

const imageURL3 = 'https://cataas.com/cat/says/hell';

const UploadedImageBox = () => (
  <UploadedImage>
    <Image src={imageURL3} />
    <p>Details title</p>
  </UploadedImage>
);

export default UploadedImageBox;
