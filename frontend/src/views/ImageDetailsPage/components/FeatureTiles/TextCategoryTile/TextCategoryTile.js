import React from 'react';
import { FeatureTemplate, FeatureItem } from 'assets/styles/style';

export const TextCategoryTile = ({ categories }) => {
  return (
    <FeatureTemplate>
      <p>Text categories</p>
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
