import styled from 'styled-components';

export const EmotionsTemplate = styled.div`
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;

  p {
    text-transform: capitalize;
  }

  div {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-row-gap: 1%;
  }
`;

export const EmotionBar = styled.div`
  display: block;
  align-self: center;
  width: ${({ emotionValue }) => {
    switch (emotionValue) {
      case 5:
        return '100%';
      case 4:
        return '80%';
      case 3:
        return '60%';
      case 2:
        return '40%';
      case 1:
        return '15%';
      default:
        return '0%';
    }
  }};
  max-width: 80%;
  height: 80%;
  margin-left: 8%;
  border-radius: 5px;
  background: rgb(32, 104, 223);
  background-image: linear-gradient(244deg, rgba(0, 232, 51, 0.5), rgba(60, 144, 228, 0));
`;
