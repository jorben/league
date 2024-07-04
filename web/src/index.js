import React from 'react';
import ReactDOM from 'react-dom/client';
import {BrowserRouter as Router, Routes, Route, Navigate} from 'react-router-dom'
import {mainRoutes} from "./routes"
import AdminFrame from "./pages/admin/layout/AdminFrame"

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/admin/*" element={<AdminFrame/>} />
        {mainRoutes.map((route, index) => {
          return <Route key={index} {...route} />;
        })}
        <Route path="*" element={<Navigate to="/404" replace />} />
      </Routes>
    </Router>
  </React.StrictMode>
);

