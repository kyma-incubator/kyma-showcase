import ImageTile from 'views/HomePage/components/ImageTile/ImageTile';
import { UploadedImagesSection } from 'views/HomePage/components/ImagesSection/ImagesSection.styles';
import { APIGET } from 'API';
import { useState, useEffect } from 'react';

const ImagesSection = () => {
  const [images, setImages] = useState([]);
  const [errorMessage, setErrorMessage] = useState('')
  const API_URL = 'http://localhost:8081/v1/images';

  const callAPIGet = async () => {
    try {
      setImages(await APIGET(API_URL));
    } catch (err) {
      console.error(err);
      setErrorMessage('Something went wrong.')
    }
  };

  useEffect(() => {
    callAPIGet();
  }, []);

  return (
    <>
      <p>{errorMessage}</p>
    <UploadedImagesSection>
      {images && images.map(({ url,img }) => {
        return <ImageTile url={url} img={img} key={img} />;
      })}
    </UploadedImagesSection>
  </>
  );
};

export default ImagesSection;
