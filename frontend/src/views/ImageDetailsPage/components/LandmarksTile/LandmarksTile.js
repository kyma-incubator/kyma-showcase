import React from 'react';
import { Template, LandmarkTitle } from './LandmarksTile.styles';

export const LandmarksTile = ({ landmarks }) => {
  return (
    <Template>
      <p>Landmarks</p>
      <ul>
        {landmarks?.map((obj, i) => (
          <li key={i}>
            <LandmarkTitle>{obj.name}</LandmarkTitle>
            <p>Latitude: {obj.latitude}</p>
            <p>Longitude: {obj.longitude}</p>
          </li>
        ))}
      </ul>
    </Template>
  );
};
