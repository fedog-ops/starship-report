import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Captain from './components/Captain';
import Officer from './components/Officer';
import Crew from './components/Crew';

const App = () => {
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
          {/* <Routes>
            
            <Route path="/captain" component={Captain} />
            <Route path="/officer" component={Officer} />
            <Route path="/crew" component={Crew} />
            <Route path="/" exact>
              {/* <h2>Select an option from the menu</h2> */}
            {/* </Route>
          </Routes> */} 
        </div>
      </div>
    </Router>
  );
};

export default App;
