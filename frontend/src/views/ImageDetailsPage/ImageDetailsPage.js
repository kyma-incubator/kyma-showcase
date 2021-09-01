import React from 'react';
import { Wrapper, Button } from 'assets/styles/style';
import Header from 'components/Header/Header';
import Footer from 'components/Footer/Footer';
import ImageDetailsArea from 'views/ImageDetailsPage/components/ImageDetailsArea/ImageDetailsArea';
import ImageDetails from 'views/ImageDetailsPage/components/ImageDetails/ImageDetails';
import { Link } from 'react-router-dom';
import { useParams } from 'react-router';
import { DetailsContextProvider } from 'contexts/detailsContext';

const ImageDetailsPage = () => {
  const { id } = useParams();
  return (
    <DetailsContextProvider>
      <Wrapper>
        <Header />
        <ImageDetailsArea id={id} />
        <ImageDetails />
        <Link to="/">
          <Button>Home Page</Button>
        </Link>
      </Wrapper>
      <Footer />
    </DetailsContextProvider>
  );
};

export default ImageDetailsPage;
