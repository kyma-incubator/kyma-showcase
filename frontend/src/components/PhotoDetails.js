import React from "react";
import styled from "styled-components";
import { GlobalStyle } from "../assets//GlobalStyle";
import Header from "./Header/Header";
import Footer from "./Footer/Footer";
import ImageDetailsArea from "./ImageDetailsArea/ImageDetailsArea";
import ImageDetails from "./ImageDetails/ImageDetails";

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  height: 100vh; //max-content
`;

const Return = styled.button`
  border-radius: 25px;
  font-size: 18px;
  font-weight: 500;
  padding: 4px 18px;
  line-height: 46px;
  transition: background-color 0.2s ease-out 0s;
  cursor: pointer;
  border: 2px solid rgb(0, 119, 225);
  background-color: rgb(0, 119, 225);
  color: rgb(255, 255, 255);
  height: 40px;
  &:hover {
    background-color: #025eb3;
    border: 2px solid #025eb3;
  }
`;

function handleClick() {
  alert("Return to home feed page");
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
