import ImageTile from 'views/HomePage/components/ImageTile/ImageTile';
import { UploadedImagesSection } from 'views/HomePage/components/ImagesSection/ImagesSection.styles';
import { APIGET } from 'API';
import { useState, useEffect } from 'react';

const ImagesSection = () => {
  const [images, setImages] = useState([]);
  const API_URL = 'http://localhost:8081/v1/images';

  const callAPIGet = async () => {
    try {
      // setImages(await APIGET(API_URL));
      console.log(await APIGET(API_URL));
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    callAPIGet();
  }, []);

  return (
    <UploadedImagesSection>
      {images.map(({ id, base64 }) => {
        return <ImageTile url={base64} id={id} key={id} />;
      })}
    </UploadedImagesSection>
  );
};

export default ImagesSection;
