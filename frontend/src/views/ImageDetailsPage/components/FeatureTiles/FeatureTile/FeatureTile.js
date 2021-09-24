import { FeatureTemplate, FeatureTitle, FeatureCarousel } from 'assets/styles/style';

export const featuresToChunks = (array) => {
  const copiedArray = [...array];
  console.log(copiedArray);

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

  console.log(result);
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
