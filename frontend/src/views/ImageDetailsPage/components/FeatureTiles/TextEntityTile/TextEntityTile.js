import React from 'react';
import { FeatureTemplate, FeatureItem } from 'assets/styles/style';

export const TextEntityTile = ({ entity }) => {
  return (
    <FeatureTemplate>
      <p>Text Entity Details</p>
      <ul>
        {entity?.map((obj, i) => (
          <FeatureItem key={i}>
            <p>{obj.name}</p>
            <p>Type: {obj.type}</p>
            <p>Sentiment: {obj.sentiment}</p>
          </FeatureItem>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
