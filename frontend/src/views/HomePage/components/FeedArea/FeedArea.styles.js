import styled from 'styled-components';

export const FeedArea = styled.main`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 30px 50px;
  border-radius: 30px;
  width: 90%;
  border: 1px solid black;
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
    font-size: 20px;
    text-transform: uppercase;
    text-align: center;
    color: white;
    letter-spacing: 3px;
  }
`;
