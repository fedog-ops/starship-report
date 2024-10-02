// src/components/Captain.js
import React, { useEffect, useState } from 'react';
import { getCaptain } from '../services/api'; // Import the getCaptain function

const Captain = () => {
  const [captain, setCaptain] = useState(null);

  useEffect(() => {
    const fetchCaptain = async () => {
      try {
        const data = await getCaptain();
        setCaptain(data);
      } catch (error) {
        console.error('Error fetching captain data:', error);
      }
    };
    fetchCaptain();
  }, []);

  if (!captain) return <h2>Loading...</h2>;

  return (
    <div>
        
      <h2>Captain Page</h2> 
      <p>Name: {captain.Name}</p>
      <p>Username: {captain.Email}</p>
      <p>password: {captain.Password}</p>
     
    </div>
  );
};

export default Captain;
