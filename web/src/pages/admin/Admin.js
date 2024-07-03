import React from 'react'
import {Navigate, Route, Routes} from 'react-router-dom'
import {adminRoutes} from '../../routes'

function Admin() {
  return (
    <div>
      <h1>Admin</h1>
      <Routes>
        {adminRoutes.map((r, index) => {
            return <Route key={index} {...r} />
        })}
        <Route path="*" element={<Navigate to="/404" replace />} />
      </Routes>
    </div>
  )
}

export default Admin
