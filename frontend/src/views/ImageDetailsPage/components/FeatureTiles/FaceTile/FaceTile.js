import React from 'react';
import { FeatureItem, FeatureTemplate, FeatureTitle } from 'assets/styles/style';
import Carousel from 'react-elastic-carousel';

const faceDetailsValues = {
  VERY_UNLIKELY: 'Very Unlikely',
  UNLIKELY: 'Unlikely',
  VERY_LIKELY: 'Very Likely',
  LIKELY: 'Likely',
  UNKNOWN: 'Unknown',
  POSSIBLE: 'Possible',
};

export const FaceTile = ({ faceDetails }) => {
  const isMany = faceDetails.length > 2 ? true : false;
  return (
    <FeatureTemplate isMany={isMany}>
      <FeatureTitle>Face details</FeatureTitle>
      <Carousel>
        {faceDetails?.map((obj, i) => (
          <FeatureItem key={i}>
            {Object.keys(obj).map((k, i) => (
              <p key={i}>
                {k}: {faceDetailsValues[obj[k]]}
              </p>
            ))}
          </FeatureItem>
        ))}
      </Carousel>
    </FeatureTemplate>
  );
};
