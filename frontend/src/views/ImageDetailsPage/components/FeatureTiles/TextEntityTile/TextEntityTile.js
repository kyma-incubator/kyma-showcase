import React from 'react';
import { FeatureTemplate, FeatureItem ,FeatureTitle} from 'assets/styles/style';

export const TextEntityTile = ({ entity }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text Entity Details</FeatureTitle>
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
