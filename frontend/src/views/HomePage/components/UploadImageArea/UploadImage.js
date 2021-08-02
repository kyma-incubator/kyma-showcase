import { StyledUploadImage } from './UploadImage.styles';
import { useState } from 'react';

const UploadImage = () => {
  const [base64Image, setBase64Image] = useState('');
  // const imageView = (base64Image) => {
  //   const IMG = `<img src=${base64Image} alt="upload img" />`;
  //   document.getElementById('cokolwiek').innerHTML = IMG;
  // };
  
  //TODO drag and drop obsługuje pliki inne niż powinien
  const uploadImage = async (event) => {
    console.log(event.target.files);
    if (event.target.files.length !== 0) {
      //walidacja error blob
      const image = event.target.files[0];
      const convertedImage = await convertImageToBase64(image);
      setBase64Image(convertedImage);
    } else {
      setBase64Image('');
    }
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
      <p>Upload an image </p>
      <input type="file" accept="image/*" onChange={(e) => uploadImage(e)} />
      <img src={base64Image} alt="upload img" />
    </StyledUploadImage>
  );
};

export default UploadImage;
