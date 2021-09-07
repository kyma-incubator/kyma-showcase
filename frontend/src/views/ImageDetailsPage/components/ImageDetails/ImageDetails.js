import { UploadedImagesSection } from './ImageDetails.styles';
import { TextTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextTile/TextTile';
import { LandmarksTile } from 'views/ImageDetailsPage/components/FeatureTiles/LandmarksTile/LandmarksTile';
import { FeatureTile } from 'views/ImageDetailsPage/components/FeatureTiles/FeatureTile/FeatureTile';
import { FaceTile } from 'views/ImageDetailsPage/components/FeatureTiles/FaceTile/FaceTile';

const name = (gcp) => {
  gcp = gcp?.map(JSON.parse);
  const labels = gcp?.find((obj) => Object.keys(obj).includes('label'))?.label;
  const textDetails = gcp?.find((obj) => Object.keys(obj).includes('font'));
  const objects = gcp?.find((obj) => Object.keys(obj).includes('objects'))?.objects;
  const logos = gcp?.find((obj) => Object.keys(obj).includes('logo'))?.logo;
  const landmarks = gcp?.find((obj) => Object.keys(obj).includes('landmarks'))?.landmarks;
  const faceDetails = gcp?.find((obj) => Object.keys(obj).includes('faceDetails'))?.faceDetails;

  return {
    labels,
    textDetails,
    objects,
    logos,
    landmarks,
    faceDetails,
  };
};
const ImageDetails = ({ gcp }) => {
  const { labels, textDetails, objects, logos, landmarks, faceDetails } = name(gcp);

  return (
    <UploadedImagesSection>
      {labels && <FeatureTile title={'Labels'} features={labels} />}
      {objects && <FeatureTile title={'Objects'} features={objects} />}
      {textDetails && <TextTile text={textDetails.font} />}
      {textDetails && <FeatureTile title={'Detected words'} features={textDetails.words} />}
      {logos && <FeatureTile title={'Logos'} features={logos} />}
      {landmarks && <LandmarksTile landmarks={landmarks} />}
      {faceDetails && <FaceTile faceDetails={faceDetails} />}
    </UploadedImagesSection>
  );
};

export default ImageDetails;
