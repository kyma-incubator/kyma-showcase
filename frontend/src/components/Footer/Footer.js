import styled from "styled-components";

const H4 = styled.h4``;

const FooterSection = styled.footer`
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  border-top: 1px solid black;
  text-align: center;
  background-color: #0e74de;
  color: white;
`;

const Footer = () => (
    <FooterSection>
      <H4>
        Procject by Raccoons
        <br />
        2021
      </H4>
    </FooterSection>
);

export default Footer;
