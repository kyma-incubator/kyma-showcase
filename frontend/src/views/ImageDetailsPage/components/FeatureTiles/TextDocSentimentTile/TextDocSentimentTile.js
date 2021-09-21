import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';

export const TextDocSentimentTile = ({ docSentiment }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Document analysis</FeatureTitle>
      <p>Language: {docSentiment.docLanguage}</p>
      <p>Sentiment: {docSentiment.docSentiment}</p>
    </FeatureTemplate>
  );
};
