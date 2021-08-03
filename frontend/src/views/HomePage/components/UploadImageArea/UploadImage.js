import { StyledUploadImage } from './UploadImage.styles';
import { useState } from 'react';

const UploadImage = () => {
  const [base64Image, setBase64Image] = useState('');
  // const imageView = (base64Image) => {
  //   const IMG = `<img src=${base64Image} alt="upload img" />`;
  //   document.getElementById('cokolwiek').innerHTML = IMG;
  // };

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
      const convertedImage = await convertImageToBase64(image);
      console.log(convertedImage);
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
      <button onClick={APIPOST}>POST</button>
      <button onClick={APIGET}>GET</button>
    </StyledUploadImage>
  );
};

export default UploadImage;
