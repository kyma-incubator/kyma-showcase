import styled from 'styled-components';

export const UploadedImage = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 100%;
  max-height: 100%;
  padding-top: 1%;
  border-radius: 5%;
  background-color: ${({ theme }) => theme.colors.secondary};
  cursor: pointer;

  p {
    text-align: center;
  }
`;

export const Image = styled.img`
  width: 70%;
  height: 70%;
  object-fit: contain;
`;
