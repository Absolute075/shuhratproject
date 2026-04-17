import { Navigate, Route, Routes } from 'react-router-dom';
import Layout from './components/Layout/Layout';
import HomePage from './pages/HomePage/HomePage';
import SimplePage from './pages/SimplePage/SimplePage';

export default function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/features" element={<SimplePage title="Features" />} />
        <Route path="/pricing" element={<SimplePage title="Pricing" />} />
        <Route path="/contact" element={<SimplePage title="Contact" />} />
        <Route path="/about" element={<SimplePage title="About" />} />
        <Route path="/privacyterms" element={<SimplePage title="Privacy / Terms" />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </Layout>
  );
}
