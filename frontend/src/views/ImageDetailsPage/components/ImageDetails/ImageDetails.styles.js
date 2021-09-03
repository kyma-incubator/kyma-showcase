import styled from "styled-components";

export const Details = styled.section`
  display: grid;
  grid-template-columns: 1fr 1fr;
  margin: 2%;
  width: 70%;
  text-align: center;

  p{
    margin: 2% 0;
  }

  ul{
    list-style:none;
  }

  li{
    margin-top: 1%;
  }
`;

export const DetailTitle = styled.p`
  font-size: 20px;
  font-weight: bold;
`