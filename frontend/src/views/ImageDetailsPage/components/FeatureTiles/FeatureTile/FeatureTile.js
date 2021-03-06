import { FeatureTemplate, FeatureTitle, FeatureCarousel } from 'assets/styles/style';

export const featuresToChunks = (array) => {
  const copiedArray = [...array];

  if (copiedArray.length < 8) {
    return {
      chunks: [copiedArray],
      isNotMany: true,
    };
  }

  const dividedLength = Math.floor(copiedArray.length / 7, 10);
  const helperIterator = 7;
  let displayedFeatures = 7;
  let result = [];

  for (let i = 0; i < dividedLength + 1; i++) {
    const featuresChunk = copiedArray.slice(i * helperIterator, displayedFeatures);
    if (featuresChunk.length) {
      result = [...result, featuresChunk];
      displayedFeatures += helperIterator;
    }
  }

  return {
    chunks: result,
    isNotMany: false,
  };
};

export const FeatureTile = ({ title, features, offDots }) => {
  const { chunks, isNotMany } = featuresToChunks(features);
  return (
    <FeatureTemplate>
      <FeatureTitle>{title}</FeatureTitle>
      <FeatureCarousel isNotMany={isNotMany} offDots={offDots}>
        {chunks.map((array, i) => (
          <ul key={i}>
            {array.map((element, i) => (
              <li key={i}>{element}</li>
            ))}
          </ul>
        ))}
      </FeatureCarousel>
    </FeatureTemplate>
  );
};
