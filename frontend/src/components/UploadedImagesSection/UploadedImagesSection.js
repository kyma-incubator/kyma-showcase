import styled from 'styled-components';
import UploadedImageBox from 'components/UploadedImage/UploadedImage';

const UploadedImagesSection = styled.section`
  display: grid;
  width: 1000px;
  max-width: 90%;
  grid-template-columns: 1fr 1fr;
  grid-row-gap: 20px;
  grid-column-gap: 20px;
  margin-bottom: 20px;
  justify-content: center;

  @media screen and (max-width: 900px) {
    grid-template-columns: 1fr;
  }
`;
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
