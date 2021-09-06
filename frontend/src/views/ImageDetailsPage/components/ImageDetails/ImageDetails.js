import { UploadedImagesSection } from './ImageDetails.styles';
import { LabelsTile } from 'views/ImageDetailsPage/components/LabelsTile/LabelsTile';
import { ObjectsTile } from '../ObjectsTile/ObjectsTile';
const ImageDetails = ({ gcp }) => {
  gcp = gcp?.map(JSON.parse);
  const labels = gcp?.find((obj) => Object.keys(obj).includes('label'))?.label || [];
  const textDetails = gcp?.find((obj) => Object.keys(obj).includes('font'));
  const objects = gcp?.find((obj) => Object.keys(obj).includes('objects'))?.objects || [];
  const logos = gcp?.find((obj) => Object.keys(obj).includes('logo'))?.logo || [];
  return (
    <UploadedImagesSection>
      {labels && <LabelsTile labels={labels} />}
      {objects && <ObjectsTile objects={objects} />}
    
    </UploadedImagesSection>
  );
};
//   return (
//     <UploadedImagesSection>
//       <DetailTitle>Detected text</DetailTitle>
//       <DetailTitle>Labels</DetailTitle>
//         <div>
//           <p>{textDetails && textDetails.font}</p>
//           <DetailTitle>Words in text:</DetailTitle>
//           <ul>
//           {textDetails && textDetails.words.map((word, i) => <li key={i}>{word}</li>)}
//           </ul>
//         </div>
//       <ul>
//         {labels.map(label => <li key={label}>{label}</li>)}
//       </ul>
//     </UploadedImagesSection>
//   );
// };

export default ImageDetails;
