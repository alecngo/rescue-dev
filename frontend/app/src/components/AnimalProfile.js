import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import defaultImage from './default.jpg'

const AnimalProfile = () => {
  const { id } = useParams();
  const [animal, setAnimal] = useState(null);

  useEffect(() => {
    fetch(`http://localhost:8080/animal/${id}`)
      .then(response => response.json())
      .then(data => setAnimal(data))
      .catch(error => console.error('Error fetching animal:', error));
  }, [id]);

  if (!animal) return <div>Loading...</div>;

  return (
    <div>
      {/* Display animal details here. Adjust according to your data structure. */}
      <h1>{animal.Name}</h1>
      <p>{animal.Description}</p>
      <img src={(animal.Photos && animal.Photos[0] && animal.Photos[0].Medium) || defaultImage} alt={animal.Name} />
      {/* ... more fields ... */}
    </div>
  );
};

export default AnimalProfile;
