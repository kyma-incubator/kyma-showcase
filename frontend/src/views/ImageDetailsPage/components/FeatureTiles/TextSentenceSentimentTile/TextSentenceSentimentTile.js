import React from 'react';
import { FeatureTemplate, FeatureItem } from 'assets/styles/style';

export const TextSentenceSentimentTile = ({ sentenceSentiment }) => {
  return (
    <FeatureTemplate>
      <p>Sentence analyze</p>
      <ul>
        {sentenceSentiment?.map((obj, i) => (
          <FeatureItem key={i}>
            <p>Sentence: {obj.sentence}</p>
            <p>Score: {obj.score}</p>
            <p>Magnitude: {obj.magnitude}</p>
          </FeatureItem>
        ))}
      </ul>
    </FeatureTemplate>
  );
};
