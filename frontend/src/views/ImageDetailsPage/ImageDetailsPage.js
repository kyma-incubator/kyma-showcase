import React from 'react';
import { Wrapper, Button } from 'assets/styles/style';
import Header from 'components/Header/Header';
import Footer from 'components/Footer/Footer';
import ImageDetailsArea from 'views/ImageDetailsPage/components/ImageDetailsArea/ImageDetailsArea';
import ImageDetails from 'views/ImageDetailsPage/components/ImageDetails/ImageDetails';
import { Link } from 'react-router-dom';
import { useParams } from 'react-router';
import { useState, useEffect } from 'react'
import { DetailsContextProvider } from 'contexts/detailsContext';
import { getImageDetailsFromAPI } from 'API';
import { Loader } from 'assets/styles/style';

const ImageDetailsPage = () => {
  const { id } = useParams();

  const [isLoading, setIsLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');
  const [imageDetails, setImageDetails] = useState(null)

  useEffect(() => {
    const callAPI = async () => {
      setIsLoading(true);
      try {
        setImageDetails(await getImageDetailsFromAPI(id))
      } catch (err) {
        setErrorMessage('Internal server error');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };
    callAPI();
  }, [])
  // eslint-disable-next-line react-hooks/exhaustive-deps
  return (
    <DetailsContextProvider>
      <Wrapper>
        <Header />
        {errorMessage && <p>{errorMessage}</p>}
        {isLoading && <Loader />}
        {!errorMessage && !isLoading && (<> <ImageDetailsArea content={imageDetails.content} />
          <ImageDetails gcp={imageDetails.gcp} /></>)}
        <Link to="/">
          <Button>Home Page</Button>
        </Link>
      </Wrapper>
      <Footer />
    </DetailsContextProvider>
  );
};

export default ImageDetailsPage;
