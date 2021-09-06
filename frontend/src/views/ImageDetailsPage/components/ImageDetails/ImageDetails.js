import { UploadedImagesSection } from './ImageDetails.styles';
import { LabelsTile } from 'views/ImageDetailsPage/components/LabelsTile/LabelsTile';
import { ObjectsTile } from '../ObjectsTile/ObjectsTile';
import { LogosTile } from '../LogosTile/LogosTile';
import { TextTile } from '../TextTile/TextTile';
import { WordsTile } from '../WordsTile/WordsTile';
import { LandmarksTile } from '../LandmarksTile/LandmarksTile';

const ImageDetails = ({ gcp }) => {
  gcp = gcp?.map(JSON.parse);
  const labels = gcp?.find((obj) => Object.keys(obj).includes('label'))?.label;
  const textDetails = gcp?.find((obj) => Object.keys(obj).includes('font'));
  const objects = gcp?.find((obj) => Object.keys(obj).includes('objects'))?.objects;
  const logos = gcp?.find((obj) => Object.keys(obj).includes('logo'))?.logo;
  const landmarks = gcp?.find((obj) => Object.keys(obj).includes('landmarks'))?.landmarks;

  return (
    <UploadedImagesSection>
      {labels && <LabelsTile labels={labels} />}
      {objects && <ObjectsTile objects={objects} />}
      {textDetails && <TextTile text={textDetails.font} />}
      {textDetails && <WordsTile words={textDetails.words} />}
      {logos && <LogosTile logos={logos} />}
      {landmarks && <LandmarksTile landmarks={landmarks} />}
    </UploadedImagesSection>
  );
};

export default ImageDetails;
