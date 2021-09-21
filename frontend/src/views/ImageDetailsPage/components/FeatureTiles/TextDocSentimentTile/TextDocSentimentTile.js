import React from 'react';
import { FeatureTemplate, FeatureTitle } from 'assets/styles/style';
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

export const TextDocSentimentTile = ({ docSentiment }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Document analysis</FeatureTitle>
      <SentenceAnalyze>
        <p>
          <span>Language</span>: {docSentiment.docLanguage.toUpperCase()}
        </p>
        <p>
          <span>Sentiment</span>: {docSentiment.docSentiment}
        </p>
      </SentenceAnalyze>
    </FeatureTemplate>
  );
};
