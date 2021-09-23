import { FeatureItem, FeatureTemplate, FeatureTitle, ModifiedCarousel } from 'assets/styles/style';

const faceDetailsValues = {
  VERY_UNLIKELY: 'Very Unlikely',
  UNLIKELY: 'Unlikely',
  VERY_LIKELY: 'Very Likely',
  LIKELY: 'Likely',
  UNKNOWN: 'Unknown',
  POSSIBLE: 'Possible',
};

export const FaceTile = ({ faceDetails }) => {
  const isNotMany = faceDetails.length > 2 ? false : true;

  return (
    <FeatureTemplate>
      <FeatureTitle>Face details</FeatureTitle>
      <ModifiedCarousel isNotMany={isNotMany}>
        {faceDetails?.map((obj, i) => (
          <FeatureItem key={i}>
            {Object.keys(obj).map((k, i) => (
              <p key={i}>
                {k}: {faceDetailsValues[obj[k]]}
              </p>
            ))}
          </FeatureItem>
        ))}
      </ModifiedCarousel>
    </FeatureTemplate>
  );
};
