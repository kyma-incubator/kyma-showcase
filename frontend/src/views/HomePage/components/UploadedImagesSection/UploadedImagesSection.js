import UploadedImageBox from 'views/HomePage/components/UploadedImage/UploadedImage';
import { UploadedImagesSection } from 'views/HomePage/components/UploadedImagesSection/UploadedImagesSection.styles';

const imageURL1 = 'https://cataas.com/cat/says/Gliwice';
const imageURL2 = 'https://cataas.com/cat/says/Warsaw';
const imageURL3 = 'https://cataas.com/cat/says/Munich';
const ImagesSection = () => (
  <UploadedImagesSection>
    <UploadedImageBox url={imageURL1} cityName={'Gliwice'} />
    <UploadedImageBox url={imageURL2} cityName={'Warsaw'} />
    <UploadedImageBox url={imageURL3} cityName={'Munich'} />
  </UploadedImagesSection>
);

export default ImagesSection;
