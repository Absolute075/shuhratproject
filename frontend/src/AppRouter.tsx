import { Navigate, Route, Routes } from 'react-router-dom'
import ApplyForm from './App'
import LandingPage from './LandingPage'
import PrivacyPolicyPage from './PrivacyPolicyPage'
import TermsConditionsPage from './TermsConditionsPage'

export default function AppRouter() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/apply" element={<ApplyForm />} />
      <Route path="/terms-conditions" element={<TermsConditionsPage />} />
      <Route path="/privacy-policy" element={<PrivacyPolicyPage />} />
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  )
}
