import styled from 'styled-components';

export const StyledUploadImage = styled.section`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
  margin: 30px 50px;
  border-radius: 30px;
  border: 1px solid black;
  width: 90%;

  h3 {
    margin-top: 1.5%;
  }

  p {
    font-weight: bold;
    color: red;
    font-size: 20px;
  }

  img {
    max-width: 80%;
    max-height: 15%;
    margin-bottom: 2%;
  }

  form {
    position: relative;
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: flex-start;
    margin: 2% 0;
    width: 50%;
    padding: 3%;
    background-color: #6797f7;
    border: 3px dashed ${({ theme }) => theme.colors.primary};
    border-radius: 3px;
    transition: 0.2s;
  }

  input {
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
`;
