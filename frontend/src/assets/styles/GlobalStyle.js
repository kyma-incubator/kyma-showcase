import { createGlobalStyle } from 'styled-components';

export const GlobalStyle = createGlobalStyle`
*{
    box-sizing: border-box;
    margin: 0;
    padding: 0;
    font-family: Poppins, Lato, San Francisco, Helvetica, sans-serif;
}

*, *::after, *::before{
    box-sizing: inherit;
}
`;
