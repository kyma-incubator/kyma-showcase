import UploadImage from 'views/HomePage/components/UploadImage/UploadImage';
import Feed from 'views/HomePage/components/FeedArea/FeedArea';
import Header from 'components/Header/Header';
import Footer from 'components/Footer/Footer';
import { Wrapper } from 'assets/styles/style';
import { ImagesContextProvider } from 'contexts/imagesContext';

const Home = () => {
  return (
    <ImagesContextProvider>
      <Wrapper>
        <Header />
        <UploadImage />
        <Feed />
      </Wrapper>
      <Footer />
    </ImagesContextProvider>
  );
};

export default Home;
