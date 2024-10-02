import React, { useEffect, useState } from 'react';
import { getCrew } from '../services/api'; // Import the getCaptain function
import CrewCard from './CrewCard';

const Crew = () => {
  const [crew, setCrew] = useState([]);

  useEffect(() => {
    const fetchCrew = async () => {
      try {
        const data = await getCrew();
        console.log(data)
        setCrew(data);
      } catch (error) {
        console.error('Error fetching captain data:', error);
      }
    };
    fetchCrew();
  }, []);
  
  if (!crew) return <h2>Loading...</h2>;

return (
  <>
        
    <h2>Officer Page</h2> 
    {crew.map(x => (
      <CrewCard crew={x} key={x.ID}/> 
    ))}
  </>
);

}

export default Crew