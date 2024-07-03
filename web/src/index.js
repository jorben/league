import React from 'react';
import ReactDOM from 'react-dom/client';
import {BrowserRouter as Router, Routes, Route, Navigate} from 'react-router-dom'
import reportWebVitals from './reportWebVitals';
import {mainRoutes} from "./routes"
import Admin from "./pages/admin/Admin"

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/admin/*" element={<Admin/>} />
        {mainRoutes.map((route, index) => {
          return <Route key={index} {...route} />;
        })}
        <Route path="*" element={<Navigate to="/404" replace />} />
      </Routes>
    </Router>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
