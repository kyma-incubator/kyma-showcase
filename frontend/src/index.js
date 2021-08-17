import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import { fetchAPIConfig } from './API';

fetchAPIConfig()
  .then(() => {
    ReactDOM.render(
      <React.StrictMode>
        <App />
      </React.StrictMode>,
      document.getElementById('root')
    );
  })
  .catch((err) => {
    alert('Cannot fetch config map.');
    console.error(err);
  });
