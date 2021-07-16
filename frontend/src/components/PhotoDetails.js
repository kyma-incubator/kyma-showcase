import React from "react";
import styled from "styled-components";
import { GlobalStyle } from "../assets//GlobalStyle";
import kymaLogo from "../img/KymaLogo.png";



const imageURL = "https://cataas.com/cat/says/work";
const kymaURL = "https://kyma-project.io/";

function handleClick() {
  alert("Return to home feed page");
}

const PhotoDetails = () => (
  <>
    <GlobalStyle />
      <header>
        <a href={kymaURL}>
          <img src={kymaLogo} alt="Kyma Logo" />
        </a>
      </header>
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
      <footer>
        <h4>
          Procject by Raccoons
          <br />
          2021
        </h4>
      </footer>
  </>
);

export default PhotoDetails;
