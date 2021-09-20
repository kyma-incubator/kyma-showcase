import React from 'react';
import { FeatureTemplate } from 'assets/styles/style';

export const TextDocSentimentTile = ({ docSentiment }) => {
  return (
    <FeatureTemplate>
      <p>Document analysis</p>
      <p>Language: {docSentiment.docLanguage}</p>
      <p>Sentiment: {docSentiment.docSentiment}</p>
    </FeatureTemplate>
  );
};
