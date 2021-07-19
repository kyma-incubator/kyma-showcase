import styled from "styled-components";
import kymaLogo from "img/KymaLogo.png";

const kymaURL = "https://kyma-project.io/";

const HeaderNav = styled.header`
  height: 50px;
  width: 100%;
  background-color: white;
`;

const Logo = styled.img`
  height: 100%;
`;

const Header = () => (
  <>
    <HeaderNav>
      <a href={kymaURL}>
        <Logo src={kymaLogo} alt="Kyma Logo" />
      </a>
    </HeaderNav>
  </>
);

export default Header;