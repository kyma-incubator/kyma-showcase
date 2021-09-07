import React from 'react';
import { LandmarkTitle } from './LandmarksTile.styles';
import { FeatureTemplate, FeatureItem } from 'assets/styles/style';

export const LandmarksTile = ({ landmarks }) => {
  return (
    <FeatureTemplate>
      <p>Landmarks</p>
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
