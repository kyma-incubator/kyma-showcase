import { UploadedImagesSection } from './ImageDetails.styles';
import { TextTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextTile/TextTile';
import { LandmarksTile } from 'views/ImageDetailsPage/components/FeatureTiles/LandmarksTile/LandmarksTile';
import { FeatureTile } from 'views/ImageDetailsPage/components/FeatureTiles/FeatureTile/FeatureTile';
import { FaceTile } from 'views/ImageDetailsPage/components/FeatureTiles/FaceTile/FaceTile';
import { TextCategoryTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextCategoryTile/TextCategoryTile';
import { TextEntityTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextEntityTile/TextEntityTile';
import { TextDocSentimentTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextDocSentimentTile/TextDocSentimentTile';
import { TextSentenceSentimentTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextSentenceSentimentTile/TextSentenceSentimentTile';

const findObject = (gcp, key) => {
  return gcp?.find((obj) => Object.keys(obj).includes(key));
};

const extractData = (gcp) => {
  gcp = gcp?.map(JSON.parse);
  const labels = findObject(gcp, 'label')?.label;
  const textDetails = findObject(gcp, 'font');
  const objects = findObject(gcp, 'objects')?.objects;
  const logos = findObject(gcp, 'logo')?.logo;
  const landmarks = findObject(gcp, 'landmarks')?.landmarks;
  const faceDetails = findObject(gcp, 'faceDetails')?.faceDetails;
  const textCategory = findObject(gcp, 'categories')?.categories;
  const textEntity = findObject(gcp, 'entityDetails')?.entityDetails;
  const textSentiment = findObject(gcp, 'docLanguage');

  return {
    labels,
    textDetails,
    objects,
    logos,
    landmarks,
    faceDetails,
    textCategory,
    textEntity,
    textSentiment,
  };
};

const ImageDetails = ({ gcp }) => {
  const { labels, textDetails, logos, landmarks, objects, faceDetails, textCategory, textEntity, textSentiment } = extractData(gcp);

  return (
    <UploadedImagesSection>
      {labels && <FeatureTile title={'Labels'} features={labels} />}
      {objects && <FeatureTile title={'Objects'} features={objects} />}
      {textDetails && <TextTile text={textDetails.font} />}
      {textDetails && <FeatureTile title={'Detected words'} features={textDetails.words} />}
      {logos && <FeatureTile title={'Logos'} features={logos} />}
      {landmarks && <LandmarksTile landmarks={landmarks} />}
      {faceDetails && <FaceTile faceDetails={faceDetails} />}
      {textEntity && <TextEntityTile entity={textEntity} />}
      {textCategory && <TextCategoryTile categories={textCategory} />}
      {textSentiment && <TextDocSentimentTile docSentiment={textSentiment} />}
      {textSentiment && <TextSentenceSentimentTile sentenceSentiment={textSentiment.sentencetDetails} />}
    </UploadedImagesSection>
  );
};

export default ImageDetails;
