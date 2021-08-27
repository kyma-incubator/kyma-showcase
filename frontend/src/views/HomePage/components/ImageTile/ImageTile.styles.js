import styled from 'styled-components';

export const Tile = styled.article`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 2%;
`;

export const Image = styled.div`
  border: 2px solid grey;
  width: 90%;
  height: 260px;
  
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
`;