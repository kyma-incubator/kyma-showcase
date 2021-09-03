import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router';
import { Link } from 'react-router-dom';

import { getImageDetailsFromAPI } from 'API';
import { Button, Loader, Wrapper } from 'assets/styles/style';
import Footer from 'components/Footer/Footer';
import Header from 'components/Header/Header';
import ImageDetails from 'views/ImageDetailsPage/components/ImageDetails/ImageDetails';
import ImageDetailsArea from 'views/ImageDetailsPage/components/ImageDetailsArea/ImageDetailsArea';

const ImageDetailsPage = () => {
  const { id } = useParams();

  const [isLoading, setIsLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');
  const [imageDetails, setImageDetails] = useState(null);
  const [analyzeLoading, setAnalyzeLoading] = useState(true);

  useEffect(() => {
    const callAPI = async () => {
      console.log("CallAPI called");
      setIsLoading(true);
      try {
        const imgDetails = await getImageDetailsFromAPI(id)
        setImageDetails(imgDetails);
        console.log(imgDetails.gcp);
        if (imgDetails.gcp?.length!==2) {
          console.log(imgDetails.gcp);
          setAnalyzeLoading(true);
          setTimeout(callAPI, 2000);
          console.log("Po set timeout");
        } else {
          setAnalyzeLoading(false);
          console.log("Analyzeloading na false");
        }
      } catch (err) {
        setErrorMessage('Internal server error');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };
    callAPI();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  
  return (
    <>
      <Wrapper>
        <Header />
        {errorMessage && <p>{errorMessage}</p>}
        {isLoading && <Loader />}
        {!errorMessage && !isLoading && (
          <>
            {' '}
            <ImageDetailsArea content={imageDetails.content} />
            <ImageDetails gcp={imageDetails.gcp} />
          </>
        )}
        <Link to="/">
          <Button>Home Page</Button>
        </Link>
      </Wrapper>
      <Footer />
    </>
  );
};

export default ImageDetailsPage;
