import styled from 'styled-components';

export const UploadedImage = styled.article`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 2%;
`;

export const Image = styled.div`
  border: 2px solid grey;
  width: 90%;
  height: 260px;

  p {
    padding: 5px;
    text-align: center;
  }
  img {
    width: 100%;
    height: 85%;
    object-fit: cover;
  }
`;