import { FeatureTemplate, FeatureTitle, FeatureCarousel } from 'assets/styles/style';
import { SentenceAnalysis } from './TextSentenceSentimentTile.styles';

export const TextSentenceSentimentTile = ({ sentenceSentiment }) => {
  const isNotMany = sentenceSentiment.length > 2 ? false : true;

  return (
    <FeatureTemplate>
      <FeatureTitle>Sentence analysis</FeatureTitle>
      <FeatureCarousel isNotMany={isNotMany}>
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
      </FeatureCarousel>
    </FeatureTemplate>
  );
};
