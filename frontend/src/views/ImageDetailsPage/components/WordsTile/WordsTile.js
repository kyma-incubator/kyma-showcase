import React from 'react';
import { Template } from './WordsTile.styles';

export const WordsTile = ({ words }) => {
  return (
    <Template>
      <p>Detected words</p>
      <ul>
        {words?.map((obj, i) => (
          <li key={i}>{obj}</li>
        ))}
      </ul>
    </Template>
  );
};
