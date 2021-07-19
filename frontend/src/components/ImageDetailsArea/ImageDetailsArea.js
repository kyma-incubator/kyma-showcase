import { ImageArea, Img, H2 } from "./ImageDetailsAreaStyle.js";

const imageURL = "https://cataas.com/cat/says/work";

const ImageDetailsArea = () => (
  <ImageArea>
    <H2>Image title</H2>
    <Img src={imageURL} alt="here will be" />
  </ImageArea>
);

export default ImageDetailsArea;
