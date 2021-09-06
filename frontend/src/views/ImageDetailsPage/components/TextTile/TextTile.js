import React from 'react';
import { Template } from './TextTile.styles';

export const TextTile = ({ text }) => {
  return (
    <Template>
      <p>Text</p>
      <p>{text}</p>
    </Template>
  );
};
