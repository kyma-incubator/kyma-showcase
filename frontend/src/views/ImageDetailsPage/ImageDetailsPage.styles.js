import styled from "styled-components";

export const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  min-height: 100vh;
  /* background-color: ${({ theme }) => theme.colors.primary}; */
`;

export const Button = styled.button`
  border-radius: 25px;
  font-size: 18px;
  font-weight: 500;
  padding: 4px 18px;
  line-height: 46px;
  transition: background-color 0.2s ease-out 0s;
  cursor: pointer;
  border: 2px solid rgb(0, 119, 225);
  background-color: rgb(0, 119, 225);
  color: rgb(255, 255, 255);
  height: 40px;
  &:hover {
    background-color: #025eb3;
    border: 2px solid #025eb3;
  }
`;
