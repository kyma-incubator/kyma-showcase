import ImageTile from 'views/HomePage/components/ImageTile/ImageTile';
import { UploadedImagesSection } from 'views/HomePage/components/ImagesSection/ImagesSection.styles';
import { Loader } from 'assets/styles/style';
import { useState, useContext, useEffect } from 'react';
import React from 'react';
import { ImagesContext } from 'contexts/imagesContext';

const ImagesSection = () => {
  const [errorMessage, setErrorMessage] = useState('');
  const [isLoading, setIsLoading] = useState(true);

  const { images, getImages } = useContext(ImagesContext);

  useEffect(() => {
    async function callAPI() {
      setIsLoading(true);
      try {
        setErrorMessage('');
        await getImages();
        setIsLoading(false);
      } catch (err) {
        setErrorMessage('Something went wrong');
        //setIsLoading(false);
      }
    }
    callAPI();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <>
      <p>{errorMessage}</p>
      {isLoading && <Loader />}
      <UploadedImagesSection>
        {images &&
          images.map(({ id, content }) => {
            return <ImageTile content={content} id={id} key={id} />;
          })}
      </UploadedImagesSection>
    </>
  );
};

export default ImagesSection;
