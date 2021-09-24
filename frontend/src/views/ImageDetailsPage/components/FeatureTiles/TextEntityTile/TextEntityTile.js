import { FeatureTemplate, FeatureItem, FeatureTitle, FeatureCarousel } from 'assets/styles/style';

export const TextEntityTile = ({ entity }) => {
  return (
    <FeatureTemplate>
      <FeatureTitle>Text Entity Details</FeatureTitle>
      <FeatureCarousel>
        {entity?.map((obj, i) => (
          <FeatureItem key={i}>
            <p>{obj.name}</p>
            <p>Type: {obj.type}</p>
            <p>Sentiment: {obj.sentiment}</p>
          </FeatureItem>
        ))}
      </FeatureCarousel>
    </FeatureTemplate>
  );
};
