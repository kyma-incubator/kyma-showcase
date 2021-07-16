import React from 'react';
import styled from 'styled-components';
import { GlobalStyle } from '../assets//GlobalStyle';

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  height: 100vh; //max-content
  background-color: #2556c6;

  @media screen and (min-width: 900px) {
    background-color: red;
  }
`;

const UploadImageArea = styled.section`
  margin: 30px 50px;
  border-radius: 30px;
  background-color: white;
  width: 90%;
  height: 30%;
`;

const FeedArea = styled.section`
  display: flex;
  flex-direction: column;
  margin: 30px 50px;
  border-radius: 30px;
  background-color: white;
  width: 90%;
  height: 60%;
`;

const FeedTitleArea = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 10%;
  background-color: #2556c6;
  border: 5px solid white;
  border-top-left-radius: 30px;
  border-top-right-radius: 30px;
`;

const FeedTitle = styled.h2`
  letter-spacing: 3px;
  font-size: 20px;
  text-transform: uppercase;
  color: white;
  text-align: center;
`;

const Root = () => (
  <>
    <GlobalStyle />
    <Wrapper>
      <UploadImageArea></UploadImageArea>
      <FeedArea>
        <FeedTitleArea>
          <FeedTitle>Feed</FeedTitle>
        </FeedTitleArea>
      </FeedArea>
    </Wrapper>
  </>
);

export default Root;
