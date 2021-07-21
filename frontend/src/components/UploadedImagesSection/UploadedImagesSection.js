import UploadedImageBox from "components/UploadedImage/UploadedImage";
import { UploadedImagesSection } from "components/UploadedImagesSection/UploadedImagesSection.styles";


const imageURL1 = "https://cataas.com/cat/says/wok";
const imageURL2 = "https://cataas.com/cat/says/h22";
const imageURL3 = "https://cataas.com/cat/says/hel";

const ImagesSection = () => (
  <UploadedImagesSection>
    <UploadedImageBox url={imageURL1} />
    <UploadedImageBox url={imageURL2} />
    <UploadedImageBox url={imageURL3} />
  </UploadedImagesSection>
);

export default ImagesSection;
