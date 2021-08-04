import ImageTile from 'views/HomePage/components/ImageTile/ImageTile';
import { UploadedImagesSection } from 'views/HomePage/components/ImagesSection/ImagesSection.styles';

const imageURL1 = 'https://cataas.com/cat/says/Gliwice';
const imageURL2 = 'https://cataas.com/cat/says/Warsaw';
const imageURL3 = 'https://cataas.com/cat/says/Munich';

const ImagesSection = () => (
  <UploadedImagesSection>
    <ImageTile url={imageURL1} cityName={'Gliwice'} />
    <ImageTile url={imageURL2} cityName={'Warsaw'} />
    <ImageTile url={imageURL3} cityName={'Munich'} />
  </UploadedImagesSection>
);

export default ImagesSection;
