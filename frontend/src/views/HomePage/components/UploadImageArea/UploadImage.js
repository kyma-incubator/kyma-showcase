import { StyledUploadImage } from './UploadImage.styles';
import { useState } from 'react';

const UploadImage = () => {
  const [base64Image, setBase64Image] = useState('');
  const [disabledButton, setDisableButton] = useState(true);
  // const imageView = (base64Image) => {
  //   const IMG = `<img src=${base64Image} alt="upload img" />`;
  //   document.getElementById('cokolwiek').innerHTML = IMG;
  // };

  const validateFormats = (shorterFormat, longerFormat) => {
    const acceptableFormats = ['.jpg', '.png', '.gif'];
    if (acceptableFormats.includes(shorterFormat) || longerFormat === '.jpeg') {
      return true;
    }
    return false;
  };

  const APIPOST = () => {
    fetch(`https://my-json-server.typicode.com/Lyczeq/images/images`, {
      method: 'POST',
      body: JSON.stringify({
        base64: `'${base64Image}'`,
      }),
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    })
      .then((res) => res.json())
      .then((data) => console.log(data));
  };

  const APIGET = () => {
    fetch('https://my-json-server.typicode.com/Lyczeq/images/images')
      .then((res) => res.json())
      .then((data) => console.log(data))
      .catch((err) => console.error(`Nie udalo sie ${err}`));
  };

  //TODO drag and drop obsługuje pliki inne niż powinien

  //TODO zmiana rozszerzenia z np pdf na jpg
  const uploadImage = async (event) => {
    console.log(event.target.files);
    if (event.target.files.length !== 0) {
      //walidacja error blob
      const image = event.target.files[0];
      const longerFormat = image.name.slice(-5);
      const shorterFormat = image.name.slice(-4);

      if (validateFormats(shorterFormat, longerFormat)) {
        const convertedImage = await convertImageToBase64(image);
        setBase64Image(convertedImage);
        setDisableButton(false);
      } else {
        alert('Zly format');
        setBase64Image('');
        setDisableButton(true);
      }
    } else {
      setBase64Image('');
      setDisableButton(true);
    }
  };

  //TODO:w jakiej sytuacji ten Promise może się nie udać
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
      <input type="file" accept="image/png, image/gif, image/jpg" onChange={(e) => uploadImage(e)} />
      <img src={base64Image} alt="upload img" />
      <button disabled={disabledButton} onClick={APIPOST}>
        POST
      </button>
      <button onClick={APIGET}>GET</button>
    </StyledUploadImage>
  );
};

export default UploadImage;
