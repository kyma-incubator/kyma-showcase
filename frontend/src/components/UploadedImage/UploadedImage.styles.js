import styled from 'styled-components';

export const UploadedImage = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 100%;
  max-height: 100%;
  padding-top: 2%;
  border-radius: 5%;
`;

export const Image = styled.div`
  border: 2px solid grey;
  border-radius: 11%;
  width: 90%;
  height: 260px;

  /* button {
    padding: 5px 10px;
    border-radius: 10%;
    border: none;
    position: relative;
    left: 300px;
    bottom: 40px;
    cursor: pointer;
  } */
  p {
    padding: 5px;
    text-align: center;
  }
  img {
    border-top-left-radius: 10%;
    border-top-right-radius: 10%;
    width: 100%;
    height: 85%;
    /* width: 70%;
  height: 70%; */
    object-fit: cover;
  }
`;
