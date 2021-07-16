import React from "react";
import styled from "styled-components";
import { GlobalStyle } from "../assets//GlobalStyle";
import kymaLogo from "../img/KymaLogo.png";

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  height: 100vh; //max-content

  @media screen and (max-width: 900px) {
    background-color: #2556c6;
  }
`;

const Header = styled.header`
  height: 50px;
  width: 100%;
  //border: 1px solid black;
  //background-color: #63a2ff;
`;

const Logo = styled.img`
  height: 100%;
  //background-color: white;
`;

const Footer = styled.footer`
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  border-top: 1px solid black;
  text-align: center;
  background-color: #0e74de;
  color: white;
`;

const imageURL = "https://cataas.com/cat/says/work";
const kymaURL = "https://kyma-project.io/";

function handleClick() {
  alert("Return to home feed page");
}

const PhotoDetails = () => (
  <>
    <GlobalStyle />
    <Wrapper>
      <Header>
        <a href={kymaURL}>
          <Logo src={kymaLogo} alt="Kyma Logo" />
        </a>
      </Header>
      <image>
        <h2>Image title</h2>
        <img src={imageURL} alt="here will be" />
      </image>
      <div>
        <ul>
          Objects
          <li>obj1</li>
        </ul>
        <ul>
          Labels
          <li>lbl1</li>
        </ul>
      </div>
      <nav>
        <button id="return" onClick={handleClick}>
          Home Page
        </button>
      </nav>
      <Footer>
        <h4>
          Procject by Raccoons
          <br />
          2021
        </h4>
      </Footer>
    </Wrapper>
  </>
);

export default PhotoDetails;
