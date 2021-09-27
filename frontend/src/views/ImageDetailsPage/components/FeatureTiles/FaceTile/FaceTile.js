import { FeatureItem, FeatureTemplate, FeatureTitle, FeatureCarousel } from 'assets/styles/style';
import EmotionsCard from './EmotionsCard';

const getEmotionsWithValues = (obj) => {
  const emotionsObject = { ...obj };

  const faceDetailsValues = {
    VERY_LIKELY: 5,
    LIKELY: 4,
    POSSIBLE: 3,
    UNLIKELY: 2,
    VERY_UNLIKELY: 1,
    UNKNOWN: 0,
  };

  const emotionNames = Object.keys(emotionsObject);
  const emotionsMap = emotionNames.map((emotion) => {
    return {
      emotion,
      emotionValue: faceDetailsValues[emotionsObject[emotion]],
    };
  });
  return emotionsMap;
};

const extractEmotions = (faceDetails) => {
  const copiedArray = [...faceDetails];

  return copiedArray?.map(getEmotionsWithValues);
};

export const FaceTile = ({ faceDetails }) => {
  const isNotMany = faceDetails.length > 2 ? false : true;
  return (
    <FeatureTemplate>
      <FeatureTitle>Face details</FeatureTitle>
      <FeatureCarousel isNotMany={isNotMany}>
        {extractEmotions(faceDetails).map((emotionsArray) => (
          <EmotionsCard emotionsArray={emotionsArray} />
        ))}
      </FeatureCarousel>
    </FeatureTemplate>
  );
};
