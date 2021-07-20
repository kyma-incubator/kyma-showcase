import styled from 'styled-components';

export const UploadedImagesSection = styled.section`
  display: grid;
  width: 1000px;
  max-width: 90%;
  grid-template-columns: 1fr 1fr;
  grid-row-gap: 20px;
  grid-column-gap: 20px;
  margin-bottom: 20px;
  justify-content: center;

  @media screen and (max-width: 600px) {
    grid-template-columns: 1fr;
  }
`;
