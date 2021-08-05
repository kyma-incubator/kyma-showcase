import ImageTile from 'views/HomePage/components/ImageTile/ImageTile';
import { UploadedImagesSection } from 'views/HomePage/components/ImagesSection/ImagesSection.styles';
import { APIGET } from 'API';
import { useState, useEffect } from 'react';

const ImagesSection = () => {
  const [images, setImages] = useState([]);

  const callAPIGet = async () => {
    try {
      setImages(await APIGET('https://my-json-server.typicode.com/Lyczeq/images/images'));
    } catch (err) {
      console.log(err);
    }
  };
  useEffect(() => {
    callAPIGet();
  }, []);
  return(
  <UploadedImagesSection>
    {images.map(({ id, base64 }) => {
        return <ImageTile url={base64} id={id} key={id} />;
      })}
  </UploadedImagesSection>
  )
};

export default ImagesSection;
