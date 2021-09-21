import React from 'react';
import { FeatureTemplate, FeatureItem, FeatureTitle } from 'assets/styles/style';
import Carousel from 'react-elastic-carousel';
import styled from 'styled-components';

const SentenceAnalyze = styled.div`
  margin: 0 20px;
  overflow-y: auto;

  p {
    line-height: 1.3;
    text-align: left;
  }

  span {
    font-weight: 600;
    font-size: 1.2rem;
    color: rgb(72, 87, 102);
  }
`;

export const TextSentenceSentimentTile = ({ sentenceSentiment }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Sentence analysis</FeatureTitle>
      <Carousel>
        {sentenceSentiment?.map((obj, i) => (
          <FeatureItem key={i}>
            <SentenceAnalyze>
              <p>
                <span>Sentence</span>: {obj.sentence}
              </p>
              <p>
                <span>Sentiment</span>: {obj.sentiment}
              </p>
            </SentenceAnalyze>
          </FeatureItem>
        ))}
      </Carousel>
    </FeatureTemplate>
  );
};
