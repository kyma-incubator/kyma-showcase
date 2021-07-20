import React from 'react';
import { GlobalStyle } from '../assets/styles/GlobalStyle';
import { Wrapper, Return } from './PhotoDetails.styles';
import Header from './Header/Header';
import Footer from './Footer/Footer';
import ImageDetailsArea from './ImageDetailsArea/ImageDetailsArea';
import ImageDetails from './ImageDetails/ImageDetails';

function handleClick() {
  alert('Return to home feed page');
}

const PhotoDetails = () => (
  <>
    <GlobalStyle />
    <Wrapper>
      <Header />
      <ImageDetailsArea />
      <ImageDetails />
      <Return id="return" onClick={handleClick}>
        Home Page
      </Return>
      <Footer />
    </Wrapper>
  </>
);

export default PhotoDetails;
