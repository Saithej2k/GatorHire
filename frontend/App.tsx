import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import HomePage from './pages/HomePage';
import JobListingsPage from './pages/JobListingsPage';
import JobDetailsPage from './pages/JobDetailsPage';
import ApplicationPage from './pages/ApplicationPage';
import ProfilePage from './pages/ProfilePage';
import NotFoundPage from './pages/NotFoundPage';
import LoginPage from './pages/LoginPage';
import SignupPage from './pages/SignupPage';
import AdminLoginPage from './pages/AdminLoginPage';
import DashboardPage from './pages/DashboardPage';
import AdminDashboardPage from './pages/AdminDashboardPage';
import CreateJobPage from './pages/CreateJobPage';
import { AuthProvider, useAuth } from './context/AuthContext';

// Protected route component
// const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
//   const { isAuthenticated } = useAuth();
  
//   if (!isAuthenticated) {
//     // Redirect to login page with the current path as the redirect parameter
//     const currentPath = window.location.pathname;
//     return <Navigate to={`/login?redirect=${encodeURIComponent(currentPath)}`} replace />;
//   }
  
//   return <>{children}</>;
// };

// Admin protected route
// const AdminRoute = ({ children }: { children: React.ReactNode }) => {
//   const { isAuthenticated, user } = useAuth();
  
//   if (!isAuthenticated) {
//     return <Navigate to="/admin/login" replace />;
//   }
  
//   if (user?.role !== 'admin') {
//     return <Navigate to="/dashboard" replace />;
//   }
  
//   return <>{children}</>;
// };

function App() {
  return (
    <AuthProvider>
      <Router>
        <div className="flex flex-col min-h-screen">
          <Navbar />
          <main className="flex-grow">
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/signup" element={<SignupPage />} />
              <Route path="/admin/login" element={<AdminLoginPage />} />
              <Route path="/jobs" element={<JobListingsPage />} />
              <Route path="/jobs/:id" element={<JobDetailsPage />} />
              <Route path="/apply/:id" element={
                // <ProtectedRoute>
                  <ApplicationPage />
                // </ProtectedRoute>
              } />
              <Route path="/profile" element={
                // <ProtectedRoute>
                  <ProfilePage />
                // </ProtectedRoute>
              } />
              <Route path="/dashboard" element={
                // <ProtectedRoute>
                  <DashboardPage />
                // </ProtectedRoute>
              } />
              <Route path="/admin/dashboard" element={
                // <AdminRoute>
                  <AdminDashboardPage />
                // </AdminRoute>
              } />
              <Route path="/admin/jobs/new" element={<CreateJobPage />} />
              <Route path="*" element={<NotFoundPage />} />
            </Routes>
          </main>
          <Footer />
        </div>
      </Router>
    </AuthProvider>
  );
}

export default App;