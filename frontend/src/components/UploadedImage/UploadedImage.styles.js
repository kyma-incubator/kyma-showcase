import styled from 'styled-components';

export const UploadedImage = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 2%;
  border-radius: 5%;
`;

export const Image = styled.div`
  border: 2px solid grey;
  border-radius: 11%;
  width: 90%;
  height: 260px;

  p {
    padding: 5px;
    text-align: center;
  }
  img {
    border-top-left-radius: 10%;
    border-top-right-radius: 10%;
    width: 100%;
    height: 85%;
    object-fit: cover;
  }
`;
