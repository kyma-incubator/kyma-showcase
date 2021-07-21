import UploadImage from 'components/UploadImageArea/UploadImage';
import Feed from "components/FeedArea/FeedArea";
import Header from "components/Header/Header";
import Footer from "components/Footer/Footer";
import { Wrapper } from './HomePage.styles';


const Home = () => (
    <Wrapper>
      <Header />
      <UploadImage />
      <Feed />
      <Footer />
    </Wrapper>
);

export default Home;