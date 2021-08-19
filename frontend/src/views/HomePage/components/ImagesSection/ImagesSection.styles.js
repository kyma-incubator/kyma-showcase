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

  p {
    text-align: center;
  }
`;

export const Loader = styled.div`
  .loader {
    position: relative;
    width: 5vh;
    height: 5vh;
    border: 0.8vh solid rgba(0, 0, 0, 0.2);
    border-left: 0.8vh solid #000000;
    border-radius: 50%;
    animation: load8 1.1s infinite linear;
    transition: opacity 0.3s;
    margin-top: 10%;
  }

  .loader--hide {
    opacity: 0;
  }

  @keyframes load8 {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
`;
