import ImagesSection from 'components/UploadedImagesSection/UploadedImagesSection';
import { FeedArea, FeedTitleArea } from 'components/FeedArea/FeedArea.styles';

const Feed = () => (
  <FeedArea>
    <FeedTitleArea>
      <h2>Feed</h2>
    </FeedTitleArea>
    <ImagesSection />
  </FeedArea>
);

export default Feed;
