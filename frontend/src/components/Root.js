import React from 'react';
import styled from 'styled-components';
import { GlobalStyle } from '../assets/GlobalStyle';
const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  min-height: 100vh;
  background-color: #2556c6;
`;

const UploadImageArea = styled.section`
  margin: 30px 50px;
  border-radius: 30px;
  background-color: white;
  width: 90%;
  height: 30vh;
`;

const FeedArea = styled.main`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 30px 50px;
  border-radius: 30px;
  background-color: white;
  width: 90%;
`;

const UploadedImages = styled.section`
  display: grid;
  width: 1000px;
  max-width: 90%;
  grid-template-columns: 1fr 1fr;
  grid-row-gap: 20px;
  grid-column-gap: 20px;
  margin-bottom: 20px;

  @media screen and (max-width: 900px) {
    grid-template-columns: 1fr;
  }
`;

const FeedTitleArea = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 1000px;
  max-width: 90%;
  height: 10%;
  margin: 10px 0;
  background-color: #2556c6;
  border-top-left-radius: 30px;
  border-top-right-radius: 30px;
`;

const FeedTitle = styled.h2`
  font-size: 20px;
  text-transform: uppercase;
  text-align: center;
  color: white;
  letter-spacing: 3px;
`;

const UploadedImage = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 300px;
  border-radius: 5%;
  background-color: #6CDFA3;
`;

const Image = styled.img`
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
`;

const A = styled.a`
  display: block;
`;

A.defaultProps = {
  href: '#',
};

Image.defaultProps = {
  src: './raccoons.jpeg',
  alt: 'szop',
};
const Root = () => (
  <>
    <GlobalStyle />
    <Wrapper>
      <UploadImageArea>
        hejka
      </UploadImageArea>
      <FeedArea>
        <FeedTitleArea>
          <FeedTitle>Feed</FeedTitle>
        </FeedTitleArea>
        <UploadedImages>
          <UploadedImage></UploadedImage> <UploadedImage></UploadedImage><UploadedImage></UploadedImage>
        </UploadedImages>
      </FeedArea>
    </Wrapper>
  </>
);

export default Root;
