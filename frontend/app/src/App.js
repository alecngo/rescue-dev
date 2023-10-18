import React, { useState } from 'react';
import AnimalList from './components/AnimalList';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import AnimalProfile from './components/AnimalProfile';

function App() {
  const [animals, setAnimals] = useState([]);
  const [zip, setZip] = useState('');
  const [distance, setDistance] = useState('');
  const DISTANCE = 50;  // Default distance

  function fetchAnimalsNearMe() {
    const actualDistance = distance ? distance : DISTANCE;
    fetch(`http://localhost:8080/findPets?zip=${zip}&distance=${actualDistance}`)
    .then(response => response.json())
    .then(data => {
        if(Array.isArray(data)) {
            setAnimals(data);
        } else {
            console.error('Unexpected API response:', data);
        }
    })
    .catch(err => {
        console.error('Error fetching animals:', err);
    });
  }

  return (
    <Router>
      <div className="container mt-5">
          <h1 className="text-center mb-5">Find Pets Near Me</h1>

          <div className="search-section mb-5">
              <div className="form-row">
                  <div className="col-md-5">
                      <label>Enter your Zip Code:</label>
                      <input
                          type="text"
                          className="form-control"
                          value={zip}
                          onChange={e => setZip(e.target.value)}
                          placeholder="Zip Code"
                      />
                  </div>
                  <div className="col-md-5">
                      <label>Enter your desired distance:</label>
                      <input
                          type="text"
                          className="form-control"
                          value={distance}
                          onChange={e => setDistance(e.target.value)}
                          placeholder="Miles"
                      />
                  </div>
                  <div className="col-md-2 d-flex align-items-end">
                      <button className="btn btn-primary w-100" onClick={fetchAnimalsNearMe}>Find</button>
                  </div>
              </div>
          </div>

          <Routes>
            <Route path="/" element={<AnimalList animals={animals} />} />
            <Route path="/animal/:id" element={<AnimalProfile />} />
          </Routes>

      </div>
    </Router>
  );  
}

export default App;
