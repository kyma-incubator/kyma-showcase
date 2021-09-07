import React from 'react';
import { FeatureTemplate } from 'assets/styles/style';

export const TextTile = ({ text }) => {
  return (
    <FeatureTemplate>
      <p>Text</p>
      <p>{text}</p>
    </FeatureTemplate>
  );
};
