import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';
import { TextWrapper } from './TextTile.styles';

export const TextTile = ({ text }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text</FeatureTitle>
      <TextWrapper>
        <p>{text}</p>
      </TextWrapper>
    </FeatureTemplate>
  );
};
