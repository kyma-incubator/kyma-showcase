import { Details } from './ImageDetails.styles';
import { useContext, useEffect, useState } from 'react';
import { DetailsContext } from 'contexts/detailsContext';

const ImageDetails = () => {
  const { imageDetails } = useContext(DetailsContext);
  const [labels, setLabels] = useState({})
  // useEffect(() => {
  //   setLabels(imageDetails.gcp.find(obj => Object.keys(obj).includes('labels')))
  // }, [])
  console.log(imageDetails)
  console.log(labels)

  return (
    <Details>
      <p>Objects</p>
      <ul>
        <li>Obj1</li>
      </ul>
      <p>Labels</p>
      <ul>
      </ul>
    </Details>
  );
};
//{labels && labels.labels.map(label => <li>{label}</li>)}

export default ImageDetails;
