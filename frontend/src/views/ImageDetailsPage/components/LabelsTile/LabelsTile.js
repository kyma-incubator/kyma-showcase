import React from 'react';
import { Template } from './LabelsTile.styles';

export const LabelsTile = ({ labels }) => {
  return (
    <Template>
      <p>Labels</p>
      <ul>
        {labels.map((label) => (
          <li key={label}>{label}</li>
        ))}
      </ul>
    </Template>
  );
};
