import React from "react";
import styled, { ThemeProvider } from "styled-components";
import { GlobalStyle } from "assets/styles/GlobalStyle";
import { theme } from "assets/styles/theme";
import { UploadImageArea } from "components/UploadImageArea/UploadImageArea";
import Feed from "components/FeedArea/FeedArea";
import Header from "components/Header/Header";
import Footer from "./Footer/Footer";

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  min-height: 100vh;
  background-color: ${theme.colors.primary};
`;

const Root = () => (
  <ThemeProvider theme={theme}>
    <GlobalStyle />
    <Wrapper>
      <Header />
      <UploadImageArea />
      <Feed></Feed>
      <Footer />
    </Wrapper>
  </ThemeProvider>
);

export default Root;