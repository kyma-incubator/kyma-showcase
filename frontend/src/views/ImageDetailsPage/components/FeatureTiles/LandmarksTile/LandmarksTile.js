import { LandmarkTitle } from './LandmarksTile.styles';
import { FeatureItem, FeatureTemplate, FeatureTitle, ModifiedCarousel } from 'assets/styles/style';

const getCoordinates = (cords) => {
  const latitudeDirection = cords.latitude > 0 ? 'N' : 'S';
  const longitudeDirection = cords.longitude > 0 ? 'W' : 'E';

  const latitudeValue = `${Math.abs(Number.parseFloat(cords.latitude).toFixed(5))} ${latitudeDirection}`;
  const longitudeValue = `${Math.abs(Number.parseFloat(cords.longitude).toFixed(5))} ${longitudeDirection} `;

  return `
        ${latitudeValue}, ${longitudeValue}
      `;
};

export const LandmarksTile = ({ landmarks }) => {
  const isNotMany = landmarks.length > 2 ? false : true;
  return (
    <FeatureTemplate>
      <FeatureTitle>Landmarks</FeatureTitle>
      <ModifiedCarousel isNotMany={isNotMany}>
        {landmarks?.map((obj, i) => (
          <FeatureItem key={i}>
            <LandmarkTitle>{obj.name}</LandmarkTitle>
            <p>{getCoordinates(obj)}</p>
          </FeatureItem>
        ))}
      </ModifiedCarousel>
    </FeatureTemplate>
  );
};
