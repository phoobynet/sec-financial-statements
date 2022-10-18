import { Route, Routes } from 'react-router-dom'
import Dashboard from './routes/dashboard/Dashboard'
import Company from './routes/company/Company'

export default function App () {
  return (
    <Routes>
      <Route
        index
        element={<Dashboard />}
      ></Route>
      <Route
        path={'/:symbol'}
        element={<Company />}
      ></Route>
    </Routes>
  )
}