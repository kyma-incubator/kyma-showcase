import { createGlobalStyle } from 'styled-components';

export const GlobalStyle = createGlobalStyle`
html{
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

background-color: red;

*, *::after, *::before{
    box-sizing: inherit;
}
`;
