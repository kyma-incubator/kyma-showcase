import styled from "styled-components";

const Details = styled.div`
  width: 70%;
  padding: 0;
  margin: 0;
  text-align: center;
  display: grid;
  grid-template-columns: 1fr 1fr;
`;

const Ul = styled.ul`
  list-style: none;
`;

const Li = styled.li``;

const ImageDetails = () => (
  <Details>
    <Ul>
      Objects
      <Li>obj1</Li>
    </Ul>
    <Ul>
      Labels
      <Li>lbl1</Li>
    </Ul>
  </Details>
);

export default ImageDetails;
