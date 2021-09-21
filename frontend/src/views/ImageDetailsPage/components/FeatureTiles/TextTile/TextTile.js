import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';
import { Text } from './TextTile.styles';

export const TextTile = ({ text }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text</FeatureTitle>
      <Text>{text}</Text>
    </FeatureTemplate>
  );
};
