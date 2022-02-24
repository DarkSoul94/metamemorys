import React from "react";
import Container from "react-bootstrap/esm/Container";
import { BrowserRouter } from "react-router-dom";
import AppRouter from "./AppRouter";
import NavBar from "./components/NavBar";

function App(){
  return (
    <BrowserRouter>
      <Container style={{width: 1000}}>
        <NavBar/>
        <AppRouter/>
      </Container>
    </BrowserRouter>
  );  
}

export default App