import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Captain from './components/Captain';
import Officer from './components/Officer';
import Crew from './components/Crew';

const App = () => {
  
  return (
    <Router>
      <nav>
        {/* Links to navigate between pages */}
        <ul>
          <li><Link to="/">Home</Link></li>
          <li><Link to="/captain">About</Link></li>
          <li><Link to="/officer">Offcier</Link></li>
          <li><Link to="/crew">Crew</Link></li>
        </ul>
      </nav>

      <Routes>
        {/* Define routes for the three pages */}
        <Route path="/" element={<p>Welcome aboard (&#128640;) Logged in as admin</p>} />
        <Route path="/captain" element={<Captain />} />
        <Route path="/officer" element={<Officer />} />
        <Route path="/crew" element={<Crew />} />
      </Routes>
    </Router>
  );
  
  return (
    <Router>
      <div style={{ display: 'flex' }}>
        <nav style={{ width: '200px', padding: '10px', borderRight: '10px solid #ccc' }}>
          <ul style={{ listStyleType: 'none', padding: 0 }}>
            <li>
              <Link to="/captain">Captain</Link>
            </li>
            <li>
              <Link to="/officer">Officer</Link>
            </li>
            <li>
              <Link to="/crew">Crew</Link>
            </li>
          </ul>
        </nav>
        <div style={{ padding: '20px', flexGrow: 1 }}>
          <Captain/>
          <Officer/>
          <Crew/>
          <Routes>
            
            <Route path="/captain" component={Captain} />
            <Route path="/officer" component={Officer} />
          </Routes>
        </div>
      </div>
    </Router>
  );
};

export default App;
