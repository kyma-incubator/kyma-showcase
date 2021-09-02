import { Details } from './ImageDetails.styles';

const ImageDetails = ({ gcp }) => {
  gcp = gcp.map(JSON.parse)
  const labels = gcp?.find(obj => Object.keys(obj).includes('label'))?.label || [];

  return (
    <Details>
      <p>Objects</p>
      <p>Labels</p>
      <ul>
        <li>Obj1</li>
      </ul>
      <ul>
        {labels.map(label => <li>{label}</li>)}
      </ul>
    </Details>
  );
};
//{labels && labels.labels.map(label => <li>{label}</li>)}

export default ImageDetails;
