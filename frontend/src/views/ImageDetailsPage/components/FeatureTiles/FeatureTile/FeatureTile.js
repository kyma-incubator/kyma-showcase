import { FeatureTemplate, FeatureTitle, FeatureCarousel } from 'assets/styles/style';

export const featuresToChunks = (array) => {
  const coppiedArray = [...array];
  if (coppiedArray.length < 8) {
    return {
      chunks: [coppiedArray],
      isNotMany: true,
    };
  }

  const dividedLength = parseInt(coppiedArray.length / 7, 10);
  const helperIterator = 8;
  let displayedFeatures = 7;
  let result = [];

  for (let i = 0; i < dividedLength + 1; i++) {
    const featuresChunk = coppiedArray.slice(i * helperIterator, displayedFeatures);
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
        {chunks.map((array) => (
          <div>
            {array.map((element, i) => (
              <p key={i}>{element}</p>
            ))}
          </div>
        ))}
      </FeatureCarousel>
    </FeatureTemplate>
  );
};
