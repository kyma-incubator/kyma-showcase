import React from 'react';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';

const faceDetailsValues = {
  VERY_UNLIKELY: 'Very Unlikely',
  UNLIKELY: 'Unlikely',
  VERY_LIKELY: 'Very Likely',
  LIKELY: 'Likely',
  UNKNOWN: 'Unknown',
  POSSIBLE: 'Possible',
};

export const FaceTile = ({ faceDetails }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Face details</FeatureTitle>
      <ul>
        {faceDetails?.map((obj, i) => (
          <FeatureItem key={i}>
            {Object.keys(obj).map((k, i) => (
              <p key={i}>
                {k}: {faceDetailsValues[obj[k]]}
              </p>
            ))}
          </FeatureItem>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
