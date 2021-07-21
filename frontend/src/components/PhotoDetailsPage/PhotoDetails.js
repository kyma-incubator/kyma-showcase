import React from "react";

import { Wrapper, Return } from "./PhotoDetails.styles";
import Header from "components/Header/Header";
import Footer from "components/Footer/Footer";
import ImageDetailsArea from "components/ImageDetailsArea/ImageDetailsArea";
import ImageDetails from "components/ImageDetails/ImageDetails";
import { Link } from "react-router-dom";

const PhotoDetails = () => (
  <>
    <Wrapper>
      <Header />
      <ImageDetailsArea />
      <ImageDetails />
      <Link to="/">
        <Return id="return">Home Page</Return>
      </Link>
      <Footer />
    </Wrapper>
  </>
);

export default PhotoDetails;
