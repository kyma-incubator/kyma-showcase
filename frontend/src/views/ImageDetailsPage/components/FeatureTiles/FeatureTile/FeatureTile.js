import React from 'react';
import { FeatureTemplate } from 'assets/styles/style';

export const FeatureTile = ({ title, features }) => {
  return (
    <FeatureTemplate>
      <p>{title}</p>
      <ul>
        {features?.map((value, index) => (
          <li key={index}>{value}</li>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
