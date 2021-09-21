import React from 'react';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';
import Carousel from 'react-elastic-carousel';
export const TextEntityTile = ({ entity }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text Entity Details</FeatureTitle>
      <Carousel>
        {entity?.map((obj, i) => (
          <FeatureItem key={i}>
            <p>{obj.name}</p>
            <p>Type: {obj.type}</p>
            <p>Sentiment: {obj.sentiment}</p>
          </FeatureItem>
        ))}
      </Carousel>
    </FeatureTemplate>
  );
};
