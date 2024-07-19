import { Link } from 'react-router-dom';
import { useTheme } from '../ThemeContext/ThemeContext';

const Navbar = () => {
  const { theme, toggleTheme } = useTheme();

  return (
    <nav className={`navbar ${theme}`}>
      <ul>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/hooks">Hooks</Link></li>
        <li><Link to="/crud">CRUD</Link></li>
        <li><Link to="/axios">Axios</Link></li>
        <li><Link to="/context">Context</Link></li>
      </ul>
      <button onClick={toggleTheme}>
        {theme === 'light' ? 'Dark' : 'Light'} Mode
      </button>
    </nav>
  );
};

export default Navbar;
