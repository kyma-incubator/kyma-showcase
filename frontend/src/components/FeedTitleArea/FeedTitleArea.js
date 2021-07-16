import styled from 'styled-components';

const FeedTitleText = styled.h2`
  font-size: 20px;
  text-transform: uppercase;
  text-align: center;
  color: ${({ theme }) => theme.colors.white};
  letter-spacing: 3px;
`;

const FeedTitleArea = styled.div`
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
`;

const FeedTitle = () => (
  <FeedTitleArea>
    <FeedTitleText>Feed</FeedTitleText>
  </FeedTitleArea>
);
export default FeedTitle;
