import { UploadedImagesSection } from './ImageDetails.styles';
import { TextTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextTile/TextTile';
import { LandmarksTile } from 'views/ImageDetailsPage/components/FeatureTiles/LandmarksTile/LandmarksTile';
import { FeatureTile } from 'views/ImageDetailsPage/components/FeatureTiles/FeatureTile/FeatureTile';
import { FaceTile } from 'views/ImageDetailsPage/components/FeatureTiles/FaceTile/FaceTile';
import { TextCategoryTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextCategoryTile/TextCategoryTile';
import { TextEntityTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextEntityTile/TextEntityTile';
import { TextSentimentTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextSentimentTile/TextSentimentTile';

const name = (gcp) => {
  gcp = gcp?.map(JSON.parse);
  const labels = gcp?.find((obj) => Object.keys(obj).includes('label'))?.label;
  const textDetails = gcp?.find((obj) => Object.keys(obj).includes('font'));
  const objects = gcp?.find((obj) => Object.keys(obj).includes('objects'))?.objects;
  const logos = gcp?.find((obj) => Object.keys(obj).includes('logo'))?.logo;
  const landmarks = gcp?.find((obj) => Object.keys(obj).includes('landmarks'))?.landmarks;
  const faceDetails = gcp?.find((obj) => Object.keys(obj).includes('faceDetails'))?.faceDetails;
  const textCategory = gcp?.find((obj) => Object.keys(obj).includes('categories'))?.categories;
  const textEntity = gcp?.find((obj) => Object.keys(obj).includes('entityDetails'))?.entityDetails;
  const textSentiment = gcp?.find((obj) => Object.keys(obj).includes('docLanguage'));

  return {
    labels,
    textDetails,
    objects,
    logos,
    landmarks,
    faceDetails,
    textCategory,
    textEntity,
    textSentiment
  };
};

const ImageDetails = ({ gcp }) => {
  const { labels, textDetails, logos, landmarks, objects, faceDetails, textCategory, textEntity, textSentiment} = name(gcp);

  return (
    <UploadedImagesSection>
      {labels && <FeatureTile title={'Labels'} features={labels} />}
      {objects && <FeatureTile title={'Objects'} features={objects} />}
      {textDetails && <TextTile text={textDetails.font} />}
      {textDetails && <FeatureTile title={'Detected words'} features={textDetails.words} />}
      {logos && <FeatureTile title={'Logos'} features={logos} />}
      {landmarks && <LandmarksTile landmarks={landmarks} />}
      {faceDetails && <FaceTile faceDetails={faceDetails} />}
      {textCategory && <TextCategoryTile categories={textCategory} />}
      {textEntity && <TextEntityTile textEntity={textEntity} />}
      {textSentiment && <TextSentimentTile textSentiment={textSentiment} />}
    </UploadedImagesSection>
  );
};

export default ImageDetails;
