import {
  UploadedImage,
  Image,
} from "components/UploadedImage/UploadedImage.styles";
import { Link } from "react-router-dom";

const UploadedImageBox = ({ url }) => (
  <UploadedImage>
    <Image>
    <Link to="/ImageDetails"><img src={url} alt="cat" /></Link>
      <p>Photo details</p>
    </Image>
  </UploadedImage>
);

export default UploadedImageBox;
