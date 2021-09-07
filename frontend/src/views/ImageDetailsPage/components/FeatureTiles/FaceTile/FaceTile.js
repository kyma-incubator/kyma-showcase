import React from 'react';
import { Template } from './FaceTile.styles';

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
    <Template>
      <p>Face details</p>
      <ul>
        {faceDetails?.map((obj, i) => (
          <li key={i}>
            {Object.keys(obj).map((k, i) => (
              <p key={i}>
                {k}: {faceDetailsValues[obj[k]]}
              </p>
            ))}
          </li>
        ))}
      </ul>
    </Template>
  );
};
