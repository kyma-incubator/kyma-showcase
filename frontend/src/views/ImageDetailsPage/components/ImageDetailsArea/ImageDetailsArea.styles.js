import styled from 'styled-components';

export const ImageArea = styled.div`
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin-top: 10%;
  max-height: 50%;

  p {
    position: absolute;
    text-align: center;
    font-weight: bold;
    text-transform: uppercase;
    letter-spacing: 2px;
    font-size: 40px;
    z-index: 2;
  }

  img {
    height: 50%;
    max-width: 80%;
    filter: ${({ nsfw }) => (nsfw ? 'blur(14px)' : 'none')};
  }
`;
