import { FooterSection, Logo } from './Footer.styles';

const kymaURL = 'https://kyma-project.io/';

const Footer = () => (
  <FooterSection>
    <a href={kymaURL}>
      <Logo src={`${process.env.PUBLIC_URL}/img/KymaLogo.png`} alt="Kyma Logo" />
    </a>
    <h4>
      Procject by Raccoons
      <br />
      2021
    </h4>
  </FooterSection>
);

export default Footer;
