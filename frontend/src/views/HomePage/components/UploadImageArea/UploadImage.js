import { StyledUploadImage } from './UploadImage.styles';
import { useState } from 'react';

const UploadImage = () => {
  const [base64Image, setBase64Image] = useState('');

  const uploadImage = async (event) => {
    const image = event.target.files[0];

    const convertedImage = await convertImageToBase64(image);
    setBase64Image(convertedImage);
  };

  const convertImageToBase64 = (image) => {
    return new Promise((resolve, reject) => {
      const fileReader = new FileReader();
      fileReader.readAsDataURL(image);

      fileReader.onload = () => {
        return resolve(fileReader.result);
      };

      fileReader.onerror = (error) => {
        reject(error);
      };
    });
  };

  return (
    <StyledUploadImage>
      <p>Upload an image or drag and drop</p>
      <input type="file" accept="image/*" text="Upload file" onChange={(e) => uploadImage(e)} />
      <img src={base64Image} alt="uploaded image" />
    </StyledUploadImage>
  );
};

export default UploadImage;
