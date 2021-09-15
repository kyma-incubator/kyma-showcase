import { StyledUploadImage } from './UploadImage.styles';
import { useState, useContext, useRef } from 'react';
import { postImageToAPI } from 'API';
import { ImagesContext } from 'contexts/imagesContext';
import { Button } from 'assets/styles/style';
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
  const [disabledPost, setDisablePost] = useState(true);
  const [disabledUpload, setDidableUpload] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');
  const { getImages } = useContext(ImagesContext);
  const [fileName, setFileName] = useState('');
  const inputRef = useRef(null);

  const callAPIPost = async () => {
    setDisablePost(true);
    inputRef.current.value = null;
    try {
      await postImageToAPI(base64Image);
      await getImages();
    } catch (err) {
      console.error(err);
      setErrorMessage('Something went wrong');
    }
  };

  const handleImageUpload = async (event) => {
    if (event.target.files.length !== 0) {
      const image = event.target.files[0];
      try {
        const extension = createExtension(image);
        const size = image.size;
        const name = image.name;
        validateFile(extension, size);
        const convertedImage = await convertImageToBase64(image);
        setBase64Image(convertedImage);
        setDisablePost(false);
        setErrorMessage('');
        setFileName(name);
      } catch (err) {
        setErrorMessage(err.message);
        setBase64Image('');
        setDisablePost(true);
      }
    } else {
      setFileName('');
      setErrorMessage('');
      setBase64Image('');
      setDisablePost(true);
    }
  };

  return (
    <>
      <StyledUploadImage>
        <h3>Upload an image </h3>
        <h5>Acceptable files: png, gif, jpg</h5>
        <nav>
          <Button disabled={disabledUpload} className="upload-image" onClick={() => setDidableUpload(true)}>Upload image</Button>
          <Button disabled={!disabledUpload} className="upload-url" onClick={() => setDidableUpload(false)}>Upload URL</Button>
        </nav>
        {disabledUpload && <form className="file-form">
          <p className="file-message">Choose a file or drag and drop</p>
          {fileName && <p className="file-name">{fileName}</p>}
          <input ref={inputRef} size={0} className="file-input" type="file" accept="image/png, image/gif, image/jpg" onChange={handleImageUpload} />
        </form>}
        {!disabledUpload && <form className="url-form">
          <label>Paste image URL: 
          <input type="text" />
          </label>
        </form>
        }
        {base64Image && <img src={base64Image} alt="Chosen file" />}
        <p>{errorMessage}</p>
        <Button disabled={disabledPost} onClick={callAPIPost}>
          POST
        </Button>
      </StyledUploadImage>
    </>
  );
};

export default UploadImage;
