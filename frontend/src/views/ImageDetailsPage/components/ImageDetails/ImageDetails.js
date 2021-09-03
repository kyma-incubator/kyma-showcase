import { Details , DetailTitle} from './ImageDetails.styles';

const ImageDetails = ({ gcp }) => {
  gcp = gcp?.map(JSON.parse)
  const labels = gcp?.find(obj => Object.keys(obj).includes('label'))?.label || [];
  const textDetails = gcp?.find(obj => Object.keys(obj).includes('font'));

  return (
    <Details>
      <DetailTitle>Detected text</DetailTitle>
      <DetailTitle>Labels</DetailTitle>
        <div>
          <p>{textDetails && textDetails.font}</p>
          <DetailTitle>Words in text:</DetailTitle>
          <ul>
          {textDetails && textDetails.words.map((word, i) => <li key={i}>{word}</li>)}
          </ul>
        </div>
      <ul>
        {labels.map(label => <li key={label}>{label}</li>)}
      </ul>
    </Details>
  );
};

export default ImageDetails;
