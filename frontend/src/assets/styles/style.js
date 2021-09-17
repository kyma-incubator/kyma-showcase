import styled from 'styled-components';

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
    cursor: default;
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

export const FeatureTemplate = styled.article`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 2%;
  max-height: 260px;
  overflow-x: hidden;
  overflow-y: auto;
`;

export const FeatureItem = styled.li`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 10% 0;
`;
