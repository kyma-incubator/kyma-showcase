import validator from 'validator';
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
  const [contentImage, setContentImage] = useState('');
  const [disabledPost, setDisablePost] = useState(true);
  const [disabledUpload, setDidableUpload] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const { getImages } = useContext(ImagesContext);
  const [fileName, setFileName] = useState('');
  const inputRef = useRef(null);

  const callAPIPost = async () => {
    setDisablePost(true);
    if (inputRef.current) inputRef.current.value = null;
    try {
      await postImageToAPI(contentImage);
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
        setContentImage(convertedImage);
        setDisablePost(false);
        setErrorMessage('');
        setFileName(name);
      } catch (err) {
        setErrorMessage(err.message);
        setContentImage('');
        setDisablePost(true);
      }
    } else {
      setFileName('');
      setErrorMessage('');
      setContentImage('');
      setDisablePost(true);
    }
  };

  const handleUrlBlur = async (event) => {
    if (event.target.value) {
      if (validator.isURL(event.target.value)) {
        setContentImage(event.target.value);
        setDisablePost(false);
        setErrorMessage('');
      } else {
        setContentImage('');
        setDisablePost(true);
        setErrorMessage('Invalid URL');
      }
    } else {
      setDisablePost(true);
    }
  };

  const handleErrorFile = () =>{
    setErrorMessage('Invalid file');
    setContentImage('');
    setDisablePost(true);
  }

  const handleImageClick = () =>{
    setDidableUpload(true)
    setErrorMessage('');
    setDisablePost(true);
  }

  const handleUrlClick = () =>{
    setDidableUpload(false)
    setErrorMessage('');
    setDisablePost(true);
  }

  return (
    <>
      <StyledUploadImage>
        <h3>Upload an image </h3>
        <h5>Acceptable files: png, gif, jpg</h5>
        <nav>
          <Button disabled={disabledUpload} className="upload-image" onClick={handleImageClick}>
            Upload file
          </Button>
          <Button disabled={!disabledUpload} className="upload-url" onClick={handleUrlClick}>
            Upload URL
          </Button>
        </nav>

        {disabledUpload && (
          <form className="file-form">
            <p className="file-message">Choose a file or drag and drop</p>
            {fileName && <p className="file-name">{fileName}</p>}
            <input ref={inputRef} size={0} className="file-input" type="file" accept="image/png, image/gif, image/jpg" onChange={handleImageUpload} />
          </form>
        )}
        {!disabledUpload && (
          <form className="url-form">
            <label for="image-url">Paste image URL: </label>
            <br />
            <input type="text" id="image-url" onBlur={handleUrlBlur} />
          </form>
        )}
        {contentImage && <img src={contentImage} alt="Chosen file" onError={handleErrorFile}/>}
        <p>{errorMessage}</p>
        <Button disabled={disabledPost} onClick={callAPIPost}>
          POST
        </Button>
      </StyledUploadImage>
    </>
  );
};

export default UploadImage;
