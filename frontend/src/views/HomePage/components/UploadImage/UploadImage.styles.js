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
  height: 30vh;

  h3 {
    margin-top: 1.5%;
  }

  p {
    font-weight: bold;
    color: red;
    font-size: 20px;
  }

  img {
    max-height: 50%;
  }

  form {
    position: relative;
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: space-evenly;
    margin: 2% 0;
    width: 50%;
    padding: 25px;
    background-color: #6797f7;
    border: 3px dashed ${({ theme }) => theme.colors.primary};
    border-radius: 3px;
    transition: 0.2s;
    cursor: pointer;
  }

  input {
    position: absolute;
    left: 0;
    top: 0;
    height: 100%;
    width: 100%;
    opacity: 0;
    cursor: pointer;
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
