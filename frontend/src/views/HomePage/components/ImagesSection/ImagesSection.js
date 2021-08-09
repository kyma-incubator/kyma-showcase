import ImageTile from 'views/HomePage/components/ImageTile/ImageTile';
import { UploadedImagesSection } from 'views/HomePage/components/ImagesSection/ImagesSection.styles';
import { useState, useContext, useEffect } from 'react';
import React from 'react';
import { ImagesContext } from 'contexts/imagesContext';

const ImagesSection = () => {
  const [errorMessage, setErrorMessage] = useState('');

  const { images, getImages } = useContext(ImagesContext);
  useEffect(() => {
    async function callAPI() {
      try {
        setErrorMessage('');
        await getImages();
      } catch (err) {
        setErrorMessage('Something went wrong');
      }
    }
    callAPI();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <>
      <p>{errorMessage}</p>
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
