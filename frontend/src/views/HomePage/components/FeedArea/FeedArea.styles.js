import styled from 'styled-components';

export const FeedArea = styled.main`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 30px 50px;
  width: 90%;

  &::after {
    content: '';
    height: 90px;
  }
`;

export const FeedTitleArea = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 1000px;
  max-width: 90%;
  height: 10%;
  margin: 10px 0;
  background-color: ${({ theme }) => theme.colors.primary};
  border-top-left-radius: 30px;
  border-top-right-radius: 30px;

  h2 {
    font-size: 1.5rem;
    font-weight: normal;
    text-align: center;
    color: white;
    letter-spacing: 3px;
  }
`;
