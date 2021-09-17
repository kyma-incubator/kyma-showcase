import styled from 'styled-components';

export const StyledUploadImage = styled.section`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
  margin: 30px 50px;
  width: 90%;
  padding-bottom: 1.5%;

  h5 {
    font-weight: normal;
    margin-top: 1%;
  }

  p {
    font-weight: bold;
    color: red;
    font-size: 20px;
  }

  img {
    max-width: 80%;
    max-height: 250px;
    margin-bottom: 2%;
  }

  .file-form {
    position: relative;
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: flex-start;
    margin: 2% 0;
    min-width: 50%;
    min-height: 60px;
    padding: 3%;
    background-color: #6797f7;
    transition: 0.2s;
    background: linear-gradient(263.44deg, rgba(37, 175, 68, 0.52) 0%, #3c90e4 47.98%, rgba(105, 106, 209, 0.85) 100%);
    border-radius: 15px;
  }

  .file-input {
    position: absolute;
    inset: 0;
    width: 100%;
    height: 100%;
    opacity: 0;
    cursor: pointer;

    &::file-selector-button {
      display: none;
    }
  }

  .file-message,
  .file-name {
    text-align: center;
    font-size: 1rem;
    font-weight: normal;
    color: white;
  }

  .file-name {
    font-style: italic;
    text-overflow: ellipsis;
    overflow: hidden;
    width: 90%;
  }

  @media screen and (max-width: 600px) {
    form {
      flex-direction: column;
    }

    .file-name {
      margin-top: 10%;
    }
  }

  .url-form {
    margin: 2% 0;
    text-align: center;
    position: relative;
    flex-direction: row;
    justify-content: flex-start;
    margin: 2% 0;
    min-width: 50%;
    min-height: 60px;
    padding: 2% 3%;
    background-color: #6797f7;
    color: white;
    transition: 0.2s;
    background: linear-gradient(263.44deg, rgba(37, 175, 68, 0.52) 0%, #3c90e4 47.98%, rgba(105, 106, 209, 0.85) 100%);
    border-radius: 15px;
  }

  #image-url {
    width: 250px;
  }
`;
