import { StyledUploadImage } from './UploadImage.styles';
import { useState } from 'react';
import { APIPOST } from 'API';

const validateFile = (extension, size) => {
  const acceptableSize = 5000000;
  const acceptableExtensions = ['.jpg', '.png', '.gif', '.jpeg'];

  if (!acceptableExtensions.includes(extension)) throw new Error('Unsupported file');

  if (size > acceptableSize) throw new Error('File is too large');
};

const convertImageToBase64 = (image) => {
  return new Promise((resolve, reject) => {
    const fileReader = new FileReader();
    fileReader.readAsDataURL(image);

    fileReader.onload = () => {
      return resolve(fileReader.result);
    };

    fileReader.onerror = () => {
      reject(new Error('Something went wrong. Please, try again :('));
    };
  });
};

export const createExtension = (file) => file.name.substr(file.name.lastIndexOf('.'));

const UploadImage = () => {
  const [base64Image, setBase64Image] = useState('');
  const [disabledButton, setDisableButton] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');
  const API_URL = 'https://my-jsonl-server.typicode.com/Lyczeq/images/images';

  const callAPIPost = async () => {
    try {
      console.log(await APIPOST(base64Image, API_URL));
    } catch (err) {
      console.error(err);
    }
  };

  const handleImageUpload = async (event) => {
    if (event.target.files.length !== 0) {
      const image = event.target.files[0];

      try {
        const extension = createExtension(image);
        const size = image.size;

        validateFile(extension, size);
        const convertedImage = await convertImageToBase64(image);
        setBase64Image(convertedImage);
        setDisableButton(false);
        setErrorMessage('');
      } catch (err) {
        setErrorMessage(err.message);
        setBase64Image('');
        setDisableButton(true);
      }
    } else {
      setErrorMessage('');
      setBase64Image('');
      setDisableButton(true);
    }
  };

  return (
    <StyledUploadImage>
      <h3>Upload an image </h3>
      <h5>Acceptable files: png, gif, jpg</h5>
      <input type="file" id="file" accept="image/png, image/gif, image/jpg" onChange={handleImageUpload} />
      {base64Image && <img src={base64Image} alt="Chosen file" />}
      <p>{errorMessage}</p>
      <button disabled={disabledButton} onClick={callAPIPost}>
        POST
      </button>
    </StyledUploadImage>
  );
};

export default UploadImage;
