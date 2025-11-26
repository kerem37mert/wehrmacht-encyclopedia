import React from 'react';
import { Link } from 'react-router-dom';

const Navigation = () => {
  return (
    <nav className="navbar">
      <div className="container">
        <Link to="/" className="navbar-brand">
          Wehrmacht Encyclopedia
        </Link>
        <ul className="navbar-menu">
          <li>
            <Link to="/" className="navbar-link">Home</Link>
          </li>
          <li>
            <Link to="/generals" className="navbar-link">Generals</Link>
          </li>
          <li>
            <Link to="/dictionary" className="navbar-link">Dictionary</Link>
          </li>
          <li>
            <Link to="/battles" className="navbar-link">Battles</Link>
          </li>
        </ul>
      </div>
    </nav>
  );
};

export default Navigation;
