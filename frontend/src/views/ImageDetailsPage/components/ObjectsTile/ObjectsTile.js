import React from 'react';
import { Template } from './ObjectsTile.styles';

export const ObjectsTile = ({ objects }) => {
  return (
    <Template>
      <p>Objects</p>
      <ul>
        {objects?.map((obj, i) => (
          <li key={i}>{obj}</li>
        ))}
      </ul>
    </Template>
  );
};
