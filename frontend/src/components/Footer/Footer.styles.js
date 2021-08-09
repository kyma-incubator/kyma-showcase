import styled from "styled-components";

export const FooterSection = styled.footer`
  position: relative;
  left: 0;
  bottom: 0;
  width: 100%;
  border-top: 1px solid black;
  text-align: center;
  background-color: ${({ theme }) => theme.colors.primary};
  color: white;
`;

export const Logo = styled.img`
  height: 30px;
  background-color: white;
  border: 2px solid white;
`;