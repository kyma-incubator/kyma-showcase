import UploadedImageBox from 'components/UploadedImage/UploadedImage';
import { UploadedImagesSection } from 'components/UploadedImagesSection/UploadedImagesSection.styles';

const imageURL1 = 'https://cataas.com/cat/says/work';
const imageURL2 = 'https://cataas.com/cat/says/hi';
const imageURL3 = 'https://cataas.com/cat/says/hell';
const ImagesSection = () => (
  <UploadedImagesSection>
    <UploadedImageBox props={imageURL1} />
    <UploadedImageBox props={imageURL2} />
    <UploadedImageBox props={imageURL3} />
  </UploadedImagesSection>
);

export default ImagesSection;
