import ImagesSection from 'views/HomePage/components/ImagesSection/ImagesSection';
import { FeedArea, FeedTitleArea } from 'views/HomePage/components/FeedArea/FeedArea.styles';

const Feed = () => (
  <FeedArea>
    <FeedTitleArea>
      <h2>Feed</h2>
    </FeedTitleArea>
    <ImagesSection />
  </FeedArea>
);

export default Feed;