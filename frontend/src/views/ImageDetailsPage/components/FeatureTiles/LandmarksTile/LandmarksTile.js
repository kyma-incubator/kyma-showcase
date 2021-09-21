import React from 'react';
import { LandmarkTitle } from './LandmarksTile.styles';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';

export const LandmarksTile = ({ landmarks }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Landmarks</FeatureTitle>
      <ul>
        {landmarks?.map((obj, i) => (
          <FeatureItem key={i}>
            <LandmarkTitle>{obj.name}</LandmarkTitle>
            <p>Latitude: {obj.latitude}</p>
            <p>Longitude: {obj.longitude}</p>
          </FeatureItem>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
