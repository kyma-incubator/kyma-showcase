import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';
import Carousel from 'react-elastic-carousel';
import { SentenceAnalysis } from './TextSentenceSentimentTile.styles';

export const TextSentenceSentimentTile = ({ sentenceSentiment }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Sentence analysis</FeatureTitle>
      <Carousel>
        {sentenceSentiment?.map((obj) => (
          <SentenceAnalysis>
            <p>
              <span>Sentence</span>: {obj.sentence}
            </p>
            <p>
              <span>Sentiment</span>: {obj.sentiment}
            </p>
          </SentenceAnalysis>
        ))}
      </Carousel>
    </FeatureTemplate>
  );
};
