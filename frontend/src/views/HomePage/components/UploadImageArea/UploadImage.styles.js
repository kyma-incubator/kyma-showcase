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
  height: 25vh;

  form{
    text-align: center;
    height: 100%;
    width: 100%;
  }

  input {
    height: 60%;
    width: 50%;
    background-color: rgba(13, 13, 13, 0.3);
    border: 3px dashed white;
    cursor: pointer;
  }

  p {
    font-weight: bold;
    color: red;
    font-size: 20px;
  }

  img {
    max-height: 200px;
  }
`;
