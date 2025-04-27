import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { fetchJobApplications, updateApplicationStatus } from '../../../frontend/services/api';
import { Application } from '../../../frontend/types/application';
import { useAuth } from '../../../frontend/context/AuthContext';
import { Navigate } from 'react-router-dom';

const JobApplicationsPage: React.FC = () => {
  const { jobId } = useParams<{ jobId: string }>();
  const [applications, setApplications] = useState<Application[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const EditJobPage = () => {
    const { isAuthenticated, user } = useAuth();
  
    if (!isAuthenticated || user?.role !== 'admin') {
      return <Navigate to="/login" />;
    }

  useEffect(() => {

    const loadApplications = async () => {
      if (!jobId) return;
      try {
        const data = await fetchJobApplications(jobId);
        setApplications(data);
      } catch (err) {
        setError('Failed to load applications. Please try again.');
      } finally {
        setLoading(false);
      }
    };
    loadApplications();
  }, [jobId]);

  const handleStatusChange = async (applicationId: string, newStatus: string) => {
    try {
      await updateApplicationStatus(applicationId, newStatus);
      setApplications((prev) =>
        prev.map((app) => (app.id === applicationId ? { ...app, status: newStatus } : app))
      );
    } catch (err) {
      alert('Failed to update status. Please try again.');
    }
  };

  if (loading) return <div>Loading applications...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="max-w-5xl mx-auto p-6">
      <h2 className="text-2xl font-bold mb-4">Applications for Job {jobId}</h2>
      {applications.length === 0 ? (
        <p>No applications found for this job.</p>
      ) : (
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Applicant Name
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Email
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Action
              </th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {applications.map((app) => (
              <tr key={app.id}>
                <td className="px-6 py-4 whitespace-nowrap">{app.fullName}</td>
                <td className="px-6 py-4 whitespace-nowrap">{app.email}</td>
                <td className="px-6 py-4 whitespace-nowrap">{app.status}</td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <select
                    value={app.status}
                    onChange={(e) => handleStatusChange(app.id, e.target.value)}
                    className="border-gray-300 rounded-md p-1"
                  >
                    <option value="pending">Pending</option>
                    <option value="reviewed">Reviewed</option>
                    <option value="interview">Interview</option>
                    <option value="accepted">Accepted</option>
                    <option value="rejected">Rejected</option>
                  </select>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default JobApplicationsPage;