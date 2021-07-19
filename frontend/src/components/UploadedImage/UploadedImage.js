import {
  UploadedImage,
  Image,
} from 'components/UploadedImage/UploadedImage.styles';

const UploadedImageBox = ({ url }) => (
  <UploadedImage>
    <Image>
      <img src={url} alt="cat"></img>
      <p>Photo details</p>
    </Image>
  </UploadedImage>
);

export default UploadedImageBox;
