import { UploadedImagesSection } from './ImageDetails.styles';
import { TextTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextTile/TextTile';
import { LandmarksTile } from 'views/ImageDetailsPage/components/FeatureTiles/LandmarksTile/LandmarksTile';
import { LandmarksMapTile } from 'views/ImageDetailsPage/components/FeatureTiles/LandmarksMapTile/LandmarksMapTile';
import { FeatureTile } from 'views/ImageDetailsPage/components/FeatureTiles/FeatureTile/FeatureTile';
import { FaceTile } from 'views/ImageDetailsPage/components/FeatureTiles/FaceTile/FaceTile';
import { TextCategoryTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextCategoryTile/TextCategoryTile';
import { TextEntityTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextEntityTile/TextEntityTile';
import { TextDocSentimentTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextDocSentimentTile/TextDocSentimentTile';
import { TextSentenceSentimentTile } from 'views/ImageDetailsPage/components/FeatureTiles/TextSentenceSentimentTile/TextSentenceSentimentTile';

const findFeature = (gcp, key) => {
  return gcp?.find((obj) => Object.keys(obj).includes(key));
};

const extractData = (gcp) => {
  gcp = gcp?.map(JSON.parse);
  const labels = findFeature(gcp, 'label')?.label;
  const textDetails = findFeature(gcp, 'font');
  const objects = findFeature(gcp, 'objects')?.objects;
  const logos = findFeature(gcp, 'logo')?.logo;
  const landmarks = findFeature(gcp, 'landmarks')?.landmarks;
  const faceDetails = findFeature(gcp, 'faceDetails')?.faceDetails;
  const textCategory = findFeature(gcp, 'categories')?.categories;
  const textEntity = findFeature(gcp, 'entityDetails')?.entityDetails;
  const textSentiment = findFeature(gcp, 'docLanguage');
  let sentenceDetails = findFeature(gcp, 'docLanguage')?.sentenceDetails;
  if (findFeature(gcp, 'docLanguage')?.sentenceDetails.length === 0) sentenceDetails = null;

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
    sentenceDetails,
  };
};

const ImageDetails = ({ gcp }) => {
  const { labels, textDetails, logos, landmarks, objects, faceDetails, textCategory, textEntity, textSentiment, sentenceDetails } = extractData(gcp);

  return (
    <UploadedImagesSection>
      {labels && <FeatureTile title={'Labels'} features={labels} />}
      {objects && <FeatureTile title={'Objects'} features={objects} />}
      {textDetails && <TextTile text={textDetails.font} />}
      {textDetails && <FeatureTile title={'Detected words'} features={textDetails.words} offDots/>}
      {logos && <FeatureTile title={'Logos'} features={logos} />}
      {landmarks && <LandmarksTile landmarks={landmarks} />}
      {landmarks && <LandmarksMapTile landmarks={landmarks} />}
      {faceDetails && <FaceTile faceDetails={faceDetails} />}
      {textEntity && <TextEntityTile entity={textEntity} />}
      {textCategory && <TextCategoryTile categories={textCategory} />}
      {textSentiment && <TextDocSentimentTile docSentiment={textSentiment} />}
      {sentenceDetails && <TextSentenceSentimentTile sentenceSentiment={sentenceDetails} />}
    </UploadedImagesSection>
  );
};

export default ImageDetails;
