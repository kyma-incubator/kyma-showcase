import React from 'react';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';

export const TextCategoryTile = ({ categories }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text categories</FeatureTitle>
      <ul>
        {categories?.map((obj, i) => (
          <FeatureItem key={i}>
            <p>Name: {obj.name}</p>
            <p>Confidence: {obj.confidence}</p>
          </FeatureItem>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
