import React from "react";
import { ThemeProvider } from "styled-components";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from "views/HomePage/HomePage";
import ImageDetails from "views/ImageDetailsPage/ImageDetailsPage";
import { GlobalStyle } from "assets/styles/GlobalStyle";
import { theme } from "assets/styles/theme";

function App() {
  return (
    <ThemeProvider theme={theme}>
      <GlobalStyle />
      <Router>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route exact path="/ImageDetails/:cityName" component={ImageDetails} />
        </Switch>
      </Router>
    </ThemeProvider>
  );
}

export default App;
