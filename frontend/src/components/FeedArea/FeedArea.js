import styled from 'styled-components';
import FeedTitle from 'components/FeedTitleArea/FeedTitleArea';
import ImagesSection from 'components/UploadedImagesSection/UploadedImagesSection';
const FeedArea = styled.main`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 30px 50px;
  border-radius: 30px;
  background-color: ${({ theme }) => theme.colors.white};
  width: 90%;
`;

const Feed = () => (
  <FeedArea>
    <FeedTitle></FeedTitle>
    <ImagesSection />
  </FeedArea>
);

export default Feed;
