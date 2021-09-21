import React from 'react';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';

export const TextSentenceSentimentTile = ({ sentenceSentiment }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Sentence analysis</FeatureTitle>
      <ul>
        {sentenceSentiment?.map((obj, i) => (
          <FeatureItem key={i}>
            <p>Sentence: {obj.sentence}</p>
            <p>Sentiment: {obj.sentiment}</p>
          </FeatureItem>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
