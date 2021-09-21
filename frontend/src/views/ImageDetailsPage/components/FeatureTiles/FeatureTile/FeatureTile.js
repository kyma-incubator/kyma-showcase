import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';
import Carousel from 'react-elastic-carousel';
import styled from 'styled-components';

export const FeatureTile = ({ title, features }) => {
  const ten = features.slice(0, 7);
  const twenty = features?.slice(8, 15);
  const thrity = features?.slice(16, 23);
  const forty = features?.slice(24, 31);

  return (
    <FeatureTemplate>
      <FeatureTitle>{title}</FeatureTitle>
      <Carousel>
        {ten?.length !== 0 && (
          <div>
            {ten.map((f) => (
              <p>{f}</p>
            ))}
          </div>
        )}
        {twenty?.length !== 0 && (
          <div>
            {twenty.map((f) => (
              <p>{f}</p>
            ))}
          </div>
        )}
        {thrity?.length !== 0 && (
          <div>
            {thrity.map((f) => (
              <p>{f}</p>
            ))}
          </div>
        )}
        {forty?.length !== 0 && (
          <div>
            {forty.map((f) => (
              <p>{f}</p>
            ))}
          </div>
        )}
      </Carousel>
    </FeatureTemplate>
  );
};
