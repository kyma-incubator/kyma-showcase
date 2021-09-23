import React from 'react';
import { FeatureTemplate, TextAnalysis, FeatureTitle } from 'assets/styles/style';

const getConfidence = (confidence) => Number.parseFloat(confidence).toFixed(4);

export const TextCategoryTile = ({ categories }) => {
  console.log(categories);
  return (
    <FeatureTemplate>
      <FeatureTitle>Text categories</FeatureTitle>
      {categories?.map((obj, i) => (
        <TextAnalysis key={i}>
          <p>
            <span>Name</span>: {obj.name}
          </p>
          <p><span>Confidence</span>: {getConfidence(obj.confidence)}</p>
        </TextAnalysis>
      ))}
    </FeatureTemplate>
  );
};
