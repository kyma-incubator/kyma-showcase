import React from 'react';
import { LandmarkTitle } from './LandmarksTile.styles';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';
import Carousel from 'react-elastic-carousel';

const getCoordinates = (cords) => {
  const latitudeDirection = cords.latitude > 0 ? 'N' : 'S';
  const longitudeDirection = cords.longitude > 0 ? 'W' : 'E';

  const latitudeValue = `${Math.abs(Number.parseFloat(cords.latitude).toFixed(5))} ${latitudeDirection}`;
  const longitudeValue = `${Math.abs(Number.parseFloat(cords.longitude).toFixed(5))} ${longitudeDirection} `;

  return (
    <p>
      {latitudeValue}, {longitudeValue}
    </p>
  );
};

export const LandmarksTile = ({ landmarks }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Landmarks</FeatureTitle>
      <Carousel>
        {landmarks?.map((obj, i) => (
          <FeatureItem key={i}>
            <LandmarkTitle>{obj.name}</LandmarkTitle>
            {getCoordinates(obj)}
          </FeatureItem>
        ))}
      </Carousel>
    </FeatureTemplate>
  );
};
