import ImageTile from 'views/HomePage/components/ImageTile/ImageTile';
import { UploadedImagesSection, Loader } from 'views/HomePage/components/ImagesSection/ImagesSection.styles';
import { useState, useContext, useEffect } from 'react';
import React from 'react';
import { ImagesContext } from 'contexts/imagesContext';

const ImagesSection = () => {
  const [errorMessage, setErrorMessage] = useState('');
  const [isLoading, setIsLoading] = useState('block');

  const { images, getImages } = useContext(ImagesContext);

  useEffect(() => {
    async function callAPI() {
      setIsLoading('block');
      try {
        setErrorMessage('');
        await getImages();
        setIsLoading('none');
      } catch (err) {
        setErrorMessage('Something went wrong');
        setIsLoading('none');
      }
    }
    callAPI();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <>
      <p>{errorMessage}</p>
      <Loader>
        <div className="loader" style={{ display: isLoading }}></div>
      </Loader>
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
