import { HeaderSection, Logo } from "./Header.styles";

const kymaURL = "https://kyma-project.io/";

const Header = () => (
  <>
    <HeaderSection>
      <a href={kymaURL}>
        <Logo src={`${process.env.PUBLIC_URL}/img/KymaLogo.png`} alt="Kyma Logo" />
      </a>
    </HeaderSection>
  </>
);
export default Header;