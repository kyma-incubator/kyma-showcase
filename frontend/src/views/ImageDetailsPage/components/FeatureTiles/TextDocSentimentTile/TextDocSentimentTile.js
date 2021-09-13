import React from 'react';
import { FeatureTemplate } from 'assets/styles/style';

export const TextDocSentimentTile = ({ docSentiment }) => {
  return (
    <FeatureTemplate>
      <p>Document analyze</p>
      <p>Language: {docSentiment.docLanguage}</p>
      <p>Score: {docSentiment.docScore}</p>
      <p>Magnitude: {docSentiment.docMagnitude}</p>
    </FeatureTemplate>
  );
};
