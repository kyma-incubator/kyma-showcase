import styled from 'styled-components';

const Image = styled.img`
  width: 70%;
  height: 70%;
  object-fit: contain;
`;

const ImageDetailsTitle = styled.p`
  text-align: center;
`;

const UploadedImage = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  padding-top: 1%;
  border-radius: 5%;
  background-color: ${({ theme }) => theme.colors.secondary};
  cursor: pointer;
`;
const imageURL3 = 'https://cataas.com/cat/says/hell';

const UploadedImageBox = ({ props }) => (
  <UploadedImage>
    <Image src={imageURL3} />
    <ImageDetailsTitle>Details title</ImageDetailsTitle>
  </UploadedImage>
);

export default UploadedImageBox;
