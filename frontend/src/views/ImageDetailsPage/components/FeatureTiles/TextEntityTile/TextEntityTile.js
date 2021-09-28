import { FeatureTemplate, FeatureTitle, FeatureCarousel, TextAnalysis } from 'assets/styles/style';

export const TextEntityTile = ({ entity }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text Entity Details</FeatureTitle>
      <FeatureCarousel>
        {entity?.map((obj, i) => (
          <TextAnalysis key={i}>
            <p>{obj.name}</p>
            <p>
              <span>Type</span>: {obj.type}
            </p>
            <p><span>Sentiment</span>: {obj.sentiment}</p>
          </TextAnalysis>
        ))}
      </FeatureCarousel>
    </FeatureTemplate>
  );
};
