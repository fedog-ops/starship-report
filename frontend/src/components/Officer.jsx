import React, { useEffect, useState } from 'react';
import { getOfficer } from '../services/api'; // Import the getCaptain function
import OfficerCard from './OfficerCard';

const Officers = () => {

  const [officer, setOfficer] = useState([]);

  useEffect(() => {
    const fetchOfficer = async () => {
      try {
        const data = await getOfficer();
        setOfficer(data);
      } catch (error) {
        console.error('Error fetching captain data:', error);
      }
    };
    fetchOfficer();
  }, []);
  
  if (!officer) return <h2>Loading...</h2>;

return (
  <>
        
    <h2>Officer Page</h2> 
    {officer.map(x => (
      <OfficerCard officer={x} key={x.ID}/>
    ))}
  </>
);

   
  
}

export default Officers