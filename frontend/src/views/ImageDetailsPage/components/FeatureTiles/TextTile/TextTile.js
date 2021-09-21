import React from 'react';
import { FeatureTemplate ,FeatureTitle} from 'assets/styles/style';

export const TextTile = ({ text }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text</FeatureTitle>
      <p>{text}</p>
    </FeatureTemplate>
  );
};
