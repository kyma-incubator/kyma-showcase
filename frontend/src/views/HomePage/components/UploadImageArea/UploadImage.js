import { StyledUploadImage } from './UploadImage.styles';
import { useState } from 'react';
import { APIGET, APIPOST } from 'API';

const UploadImage = () => {
  const [base64Image, setBase64Image] = useState('');
  const [disabledButton, setDisableButton] = useState(true);

  const [errorMessage, setErrorMessage] = useState('');

  const validateFile = (extension, size) => {
    const acceptableSize = 5000000;
    const acceptableExtensions = ['.jpg', '.png', '.gif', '.jpeg'];

    if (!acceptableExtensions.includes(extension)) throw new Error('Zly format pliku');

    if (size > acceptableSize) throw new Error('Zbyt duzy plik');
  };

  const callAPIGet = async () => {
    try {
      console.log(await APIGET());
    } catch (err) {
      console.log(err);
    }
  };

  const callAPIPost = async () => {
    try {
      console.log(await APIPOST(base64Image));
    } catch (err) {
      console.log(err);
    }
  };

  //TODO zmiana rozszerzenia z np pdf na jpg - obsluga err mess z visionAI
  const handleImageUpload = async (event) => {
    console.log(event.target.files);
    if (event.target.files.length !== 0) {
      //walidacja error blob
      const image = event.target.files[0];
      try {
        const extension = image.name.substr(image.name.lastIndexOf('.'));
        const size = image.size;

        validateFile(extension, size);
        const convertedImage = await convertImageToBase64(image);
        // console.log(convertedImage);
        setBase64Image(convertedImage);
        setDisableButton(false);
        setErrorMessage('')
      } catch (err) {
        setErrorMessage(err.message);
        setBase64Image('');
        setDisableButton(true);
      }
    } else {
      setErrorMessage('')
      setBase64Image('');
      setDisableButton(true);
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
      <h3>Upload an image </h3>
      <h5>Acceptable files: png, gif, jpg</h5>
      <input type="file" accept="image/png, image/gif, image/jpg" onChange={handleImageUpload} />
      {base64Image && <img src={base64Image} alt="zdjecie" />}
      <p>{errorMessage}</p>
      <button disabled={disabledButton} onClick={callAPIPost}>
        POST
      </button>
      <button onClick={callAPIGet}>GET</button>
    </StyledUploadImage>
  );
};

export default UploadImage;
