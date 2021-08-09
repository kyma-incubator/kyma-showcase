import styled from "styled-components";

export const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 1200px;
  max-width: 100%;
  margin: 0 auto;
  min-height: 100vh;
`;

export const Button = styled.button`
  border-radius: 25px;
  font-size: 18px;
  font-weight: 500;
  padding: 4px 15px;
  line-height: 15px;
  transition: background-color 0.2s;
  cursor: pointer;
  border: 2px solid ${({ theme }) => theme.colors.primary};
  background-color: ${({ theme }) => theme.colors.primary};
  color: rgb(255, 255, 255);
  height: 30px;
  &:hover {
    background-color: ${({ theme }) => theme.colors.secondary};
    border: 2px solid ${({ theme }) => theme.colors.secondary};
  }
`;
