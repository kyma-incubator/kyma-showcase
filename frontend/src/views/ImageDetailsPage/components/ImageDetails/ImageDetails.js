import { LabelsTile } from 'views/ImageDetailsPage/components/LabelsTile/LabelsTile';
import { FaceTile } from '../FaceTile/FaceTile';
import { LandmarksTile } from '../LandmarksTile/LandmarksTile';
import { LogosTile } from '../LogosTile/LogosTile';
import { ObjectsTile } from '../ObjectsTile/ObjectsTile';
import { TextTile } from '../TextTile/TextTile';
import { WordsTile } from '../WordsTile/WordsTile';
import { UploadedImagesSection } from './ImageDetails.styles';

const ImageDetails = ({ gcp }) => {
  gcp = gcp?.map(JSON.parse);
  const labels = gcp?.find((obj) => Object.keys(obj).includes('label'))?.label;
  const textDetails = gcp?.find((obj) => Object.keys(obj).includes('font'));
  const objects = gcp?.find((obj) => Object.keys(obj).includes('objects'))?.objects;
  const logos = gcp?.find((obj) => Object.keys(obj).includes('logo'))?.logo;
  const landmarks = gcp?.find((obj) => Object.keys(obj).includes('landmarks'))?.landmarks;
  const faceDetails = gcp?.find((obj) => Object.keys(obj).includes('faceDetails'))?.faceDetails;

  return (
    <UploadedImagesSection>
      {labels && <LabelsTile labels={labels} />}
      {objects && <ObjectsTile objects={objects} />}
      {textDetails && <TextTile text={textDetails.font} />}
      {textDetails && <WordsTile words={textDetails.words} />}
      {logos && <LogosTile logos={logos} />}
      {landmarks && <LandmarksTile landmarks={landmarks} />}
      {faceDetails && <FaceTile faceDetails={faceDetails} />}
    </UploadedImagesSection>
  );
};

export default ImageDetails;
