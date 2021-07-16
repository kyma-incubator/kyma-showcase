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

const NavBar = styled.nav`
  height: 40px;
  width: 100%;
  margin: 0;
  padding: 0;
  //border: 1px solid black;
`;

const Return = styled.button`
  border-radius: 25px;
  font-size: 18px;
  font-weight: 500;
  padding: 4px 18px;
  line-height: 46px;
  transition: background-color 0.2s ease-out 0s;
  text-decoration: none;
  cursor: pointer;
  border: 2px solid rgb(0, 119, 225);
  background-color: rgb(0, 119, 225);
  color: rgb(255, 255, 255);
  height: 40px;
  left: 46%;
  position: absolute;
  &:hover {
    background-color: #025eb3;
    border: 2px solid #025eb3;
  }
`;

const ImageArea = styled.div`
  text-align: center;
  //border: 1px solid black;
  max-height: 50%;
`;

const Img = styled.img`
  max-height: 90%;
  max-width: 100%;
`;

const Details = styled.div`
  //border: 1px solid black;
  width: 70%;
  padding: 0;
  margin: 0;
  text-align: center;
`;

const Ul = styled.ul`
  list-style: none;
`;

const Li = styled.li``;

const H2 = styled.h2``;

const H4 = styled.h4``;

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
      <ImageArea>
        <H2>Image title</H2>
        <Img src={imageURL} alt="here will be" />
      </ImageArea>
      <Details>
        <Ul>
          Objects
          <Li>obj1</Li>
        </Ul>
        <Ul>
          Labels
          <Li>lbl1</Li>
        </Ul>
      </Details>
      <NavBar>
        <Return id="return" onClick={handleClick}>
          Home Page
        </Return>
      </NavBar>
      <Footer>
        <H4>
          Procject by Raccoons
          <br />
          2021
        </H4>
      </Footer>
    </Wrapper>
  </>
);

export default PhotoDetails;
