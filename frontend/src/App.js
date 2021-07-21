import React from "react";
import { ThemeProvider } from "styled-components";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from "./components/HomePage/Home";
import PhotoDetails from "components/PhotoDetailsPage/PhotoDetails";
import { GlobalStyle } from "assets/styles/GlobalStyle";
import { theme } from "assets/styles/theme";

function App() {
  return (
    <ThemeProvider theme={theme}>
      <GlobalStyle />
      <Router>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route exact path="/ImageDetails" component={PhotoDetails} />
        </Switch>
      </Router>
    </ThemeProvider>
  );
}

export default App;
