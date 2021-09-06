import styled from 'styled-components';

export const Template = styled.article`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 2%;
  max-height: 260px;
  overflow-x: hidden;
  overflow-y: auto;

  li {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 10% 0;
  }
`;

export const LandmarkTitle = styled.p`
  font-weight: bold;
  font-size: 20px;
`;
