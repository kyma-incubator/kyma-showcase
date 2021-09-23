import { FeatureTemplate, FeatureTitle, TextAnalysis } from 'assets/styles/style';

export const TextDocSentimentTile = ({ docSentiment }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Document analysis</FeatureTitle>
      <TextAnalysis>
        <p>
          <span>Language</span>: {docSentiment.docLanguage.toUpperCase()}
        </p>
        <p>
          <span>Sentiment</span>: {docSentiment.docSentiment}
        </p>
      </TextAnalysis>
    </FeatureTemplate>
  );
};
