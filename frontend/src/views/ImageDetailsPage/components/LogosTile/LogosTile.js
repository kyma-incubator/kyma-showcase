import React from 'react';
import { Template } from './LogosTile.styles';

export const LogosTile = ({ logos }) => {
  return (
    <Template>
      <p>Objects</p>
      <ul>
        {logos.map((logo, i) => (
          <li key={i}>{logo}</li>
        ))}
      </ul>
    </Template>
  );
};
