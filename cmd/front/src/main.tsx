import React from 'react';
import ReactDOM from 'react-dom/client';
import { PrimeReactProvider } from 'primereact/api';
import 'primeicons/primeicons.css';
import "primereact/resources/themes/lara-dark-blue/theme.css";
import 'primeflex/primeflex.css'

import Routes from '@/routes/routes';

import { AppContextProvider } from './context/appContextProvider';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <AppContextProvider>
      <PrimeReactProvider >
        <Routes />
      </PrimeReactProvider>
    </AppContextProvider>
  </React.StrictMode>
);