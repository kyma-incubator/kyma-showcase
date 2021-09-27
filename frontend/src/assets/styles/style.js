import styled from 'styled-components';
import Carousel from 'react-elastic-carousel';

export const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  min-height: 100vh;
`;

export const Button = styled.button`
  border-radius: 25px;
  padding: 4px 15px;
  line-height: 15px;
  transition: background-color 0.2s;
  cursor: pointer;
  border: 2px solid ${({ theme }) => theme.colors.primary};
  background-color: ${({ theme }) => theme.colors.primary};
  color: rgb(255, 255, 255);
  height: 30px;

  &:hover {
    background-color: ${({ theme }) => theme.colors.secondary};
    border: 2px solid ${({ theme }) => theme.colors.secondary};
  }

  &:disabled,
  button[disabled] {
    border: 1px solid #999999;
    background-color: #cccccc;
    color: #666666;
    cursor: not-allowed;
  }
`;

export const UploadButton = styled.button`
  border-bottom-left-radius: 25px;
  border-top-left-radius: 25px;
  padding: 4px 15px;
  line-height: 15px;
  transition: background-color 0.2s;
  cursor: pointer;
  border: 2px solid ${({ theme }) => theme.colors.primary};
  background-color: ${({ theme }) => theme.colors.primary};
  color: rgb(255, 255, 255);
  height: 30px;

  &:hover {
    background-color: ${({ theme }) => theme.colors.secondary};
    border: 2px solid ${({ theme }) => theme.colors.secondary};
  }

  &:disabled,
  button[disabled] {
    border: 1px solid #999999;
    background-color: #cccccc;
    color: #666666;
    cursor: default;
  }
`;

export const UrlButton = styled.button`
  border-bottom-right-radius: 25px;
  border-top-right-radius: 25px;
  padding: 4px 15px;
  line-height: 15px;
  transition: background-color 0.2s;
  cursor: pointer;
  border: 2px solid ${({ theme }) => theme.colors.primary};
  background-color: ${({ theme }) => theme.colors.primary};
  color: #fff;
  height: 30px;

  &:hover {
    background-color: ${({ theme }) => theme.colors.secondary};
    border: 2px solid ${({ theme }) => theme.colors.secondary};
  }

  &:disabled,
  button[disabled] {
    border: 1px solid #999999;
    background-color: #cccccc;
    color: #666666;
    cursor: default;
  }
`;

export const Loader = styled.div`
  position: relative;
  width: 5vh;
  height: 5vh;
  border: 0.8vh solid rgba(0, 0, 0, 0.2);
  border-left: 0.8vh solid black;
  border-radius: 50%;
  animation: load 1.1s infinite linear;
  transition: opacity 0.3s;
  margin: 10% 0 2%;

  @keyframes load {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
`;

export const FeatureTitle = styled.p`
  display: block;
  border-top-left-radius: 13px;
  border-top-right-radius: 13px;
  height: 10%;
  width: 100%;
  padding: 3% 0;
  background: rgb(32, 104, 223);
  background-image: linear-gradient(244deg, rgba(0, 232, 51, 0.5), rgba(60, 144, 228, 0));
  font-size: 20px;
  color: #fff;
`;

export const FeatureTemplate = styled.article`
  display: flex;
  border: 2px solid black;
  flex-direction: column;
  border-radius: 15px;
  align-items: center;
  margin-top: 2%;
  height: 260px;
`;

export const FeatureCarousel = styled(Carousel)`
  overflow-y: auto;
  overflow-x: hidden;
  height: 90%;
  justify-content: center;

  .rec.rec-arrow {
    background-color: #1da697;
    margin: 0 2%;
    width: 38px;
    height: 38px;
    min-width: 38px;
    line-height: 38px;
    color: #fff;
    display: ${({ isNotMany }) => (isNotMany ? 'none' : 'block')};

    &:hover:enabled {
      background-color: rgb(32, 104, 223);
    }

    &:disabled {
      background-color: rgba(103, 58, 183, 0.1);
      cursor: not-allowed;
    }

    &:checked {
      background-color: rgb(32, 104, 223);
    }
  }

  .rec-pagination {
    display: ${({ offDots }) => (offDots ? 'none' : 'flex')};
  }

  .rec.rec-dot {
    display: ${({ offDots, isNotMany }) => {
      if (offDots || isNotMany) return 'none';
    }};

    &:hover {
      box-shadow: 0 0 1px 3px rgb(32, 104, 223);
    }
  }

  .rec.rec-dot_active {
    background-color: rgba(32, 104, 223, 0.5);
    box-shadow: 0 0 1px 3px rgb(32, 104, 223);
  }
`;

export const TextAnalysis = styled.div`
  display: flex;
  flex-direction: column;
  margin: 0 20px;
  overflow-y: auto;
  height: 90%;
  justify-content: center;

  p {
    line-height: 1.3;
    text-align: left;
    font-size: 1.3rem;
  }

  span {
    font-weight: 600;
    font-size: 1.4rem;
    color: rgb(72, 87, 102);
  }
`;

export const FeatureItem = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
`;
