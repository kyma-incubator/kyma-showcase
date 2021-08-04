import { HeaderSection, Logo } from './Header.styles';
import { Link } from 'react-router-dom';

const Header = () => (
  <>
    <HeaderSection>
      <Link to="/">
        <Logo src={`${process.env.PUBLIC_URL}/img/KymaLogo.png`} alt="Kyma Logo" />
      </Link>
    </HeaderSection>
  </>
);
export default Header;
