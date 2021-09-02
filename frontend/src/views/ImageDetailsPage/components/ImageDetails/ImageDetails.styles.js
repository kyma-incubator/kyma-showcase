import styled from "styled-components";

export const Details = styled.section`
  display: grid;
  grid-template-columns: 1fr 1fr;
  margin: 2%;
  width: 70%;
  text-align: center;

  p{
    font-size: 20px;
    font-weight: bold;
    margin-bottom: 2%;
  }

  ul{
    list-style:none;
  }

  li{
    margin-top: 1%;
  }
`;