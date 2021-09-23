import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';
import Carousel from 'react-elastic-carousel';

export const featuresToChunks = (array) => {
  const coppiedArray = [...array];
  if (coppiedArray.length < 7) {
    return {
      chunks: [coppiedArray],
      isNotMany: true,
    };
  }

  const dividedLength = parseInt(coppiedArray.length / 7, 10);
  const helperIterator = 8;
  let displayedFeatures = 7;
  let result = [];

  for (let i = 0; i < dividedLength + 1; i++) {
    const featuresChunk = coppiedArray.slice(i * helperIterator, displayedFeatures);
    if (featuresChunk.length) {
      result = [...result, featuresChunk];
      displayedFeatures += helperIterator;
    }
  }
  return {
    chunks: result,
    isNotMany: false,
  };
};

export const FeatureTile = ({ title, features, offDots }) => {
  const { chunks, isNotMany } = featuresToChunks(features);
  return (
    <FeatureTemplate isNotMany={isNotMany} offDots={offDots}>
      <FeatureTitle>{title}</FeatureTitle>
      <Carousel>
        {chunks.map((array) => (
          <div>
            {array.map((element, i) => (
              <p key={i}>{element}</p>
            ))}
          </div>
        ))}
      </Carousel>
    </FeatureTemplate>
  );
};
