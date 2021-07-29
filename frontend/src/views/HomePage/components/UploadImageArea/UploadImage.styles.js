import styled from 'styled-components';

export const StyledUploadImage = styled.section`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
  margin: 30px 50px;
  border-radius: 30px;
  background-color: ${({ theme }) => theme.colors.white};
  border: 1px solid black;
  
  width: 90%;
  height: 25vh;
  p {
    text-align: center;
    color: ${({ theme }) => theme.colors.black};
    font-size: 20px;
  }

  div {
    height: 60%;
    width: 50%;
    background-color: rgba(13, 13, 13, 0.3);
    border: 3px dashed white;
  }
`;
