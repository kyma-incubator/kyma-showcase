import UploadImage from 'views/HomePage/components/UploadImage/UploadImage';
import Feed from 'views/HomePage/components/FeedArea/FeedArea';
import Header from 'components/Header/Header';
import Footer from 'components/Footer/Footer';
import { Wrapper } from 'assets/styles/style';

const Home = () => {
  return (
    <>
      <Wrapper>
        <Header />
        <UploadImage />
        <Feed />
      </Wrapper>
      <Footer />
    </>
  );
};

export default Home;
