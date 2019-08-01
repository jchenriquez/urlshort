import React from 'react';
// import logo from './logo.svg';
import axios from 'axios';
import Urls from './components/urls.components';
import './App.css';
import Container from 'react-bootstrap/Container';
import * as _ from 'lodash';

console.log(process.env.PUBLIC_URL);
axios.defaults.baseURL = `${process.env.PUBLIC_URL}`;
axios.defaults.headers.post['Content-Type'] = 'application/json';

function App() {
  return (
    <div className='App'>
      <Container>
        <Urls />
      </Container>
    </div>
  );
}

export default App;
