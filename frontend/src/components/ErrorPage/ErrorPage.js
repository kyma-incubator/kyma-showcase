import { Wrapper } from 'assets/styles/style';
import Header from 'components/Header/Header';
import Footer from 'components/Footer/Footer';
import { Button } from 'assets/styles/style';
import { P, Image, H1 } from './ErrorPage.styles';
import { Link } from 'react-router-dom';

const ErrorPage = () => (
  <>
    <Wrapper>
      <Header />
      <H1>404 Page Not Found</H1>
      <P>Sorry, we can't find that page.</P>
      <Image src={`${process.env.PUBLIC_URL}/img/rac_sad.jpeg`} alt="SadRaccoon" />
      <Link to="/">
        <Button>Return to home Page</Button>
      </Link>
    </Wrapper>
    <Footer />
  </>
);

export default ErrorPage;
