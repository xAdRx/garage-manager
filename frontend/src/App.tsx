import { Navigate, Route, Routes } from 'react-router-dom'
import Header from './components/Header.tsx'
import GaragePage from './pages/GaragePage.tsx'
import OwnerPanelPage from './pages/OwnerPanelPage.tsx'

export default function App() {
  return (
    <>
      <Header />
      <main className="page">
        <Routes>
          <Route path="/" element={<GaragePage />} />
          <Route path="/owner" element={<OwnerPanelPage />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </main>
    </>
  )
}
