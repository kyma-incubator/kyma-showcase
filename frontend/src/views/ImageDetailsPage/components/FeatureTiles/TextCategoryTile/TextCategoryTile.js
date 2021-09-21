import React from 'react';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';

const getConfidence = (confidence) => Number.parseFloat(confidence).toFixed(4);

export const TextCategoryTile = ({ categories }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text categories</FeatureTitle>
      {categories?.map((obj, i) => (
        <FeatureItem key={i}>
          <p>Name: {obj.name}</p>
          <p>Confidence: {getConfidence(obj.confidence)}</p>
        </FeatureItem>
      ))}
    </FeatureTemplate>
  );
};
