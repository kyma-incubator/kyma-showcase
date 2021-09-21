import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';

export const FeatureTile = ({ title, features }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>{title}</FeatureTitle>
      <ul>
        {features?.map((value, index) => (
          <li key={index}>{value}</li>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
