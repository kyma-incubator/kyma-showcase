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
            <p>type: {obj.type}</p>
            <p>sentiment: {obj.sentiment}</p>
          </FeatureItem>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
