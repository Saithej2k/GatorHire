import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { fetchJobById, updateJob } from '../frontend/services/api';
import { Job } from '../frontend/types/job';
import { useAuth } from '../frontend/context/AuthContext';
import { Navigate } from 'react-router-dom';


const EditJobPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [job, setJob] = useState<Job | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const { isAuthenticated, user } = useAuth();
  
    if (!isAuthenticated || user?.role !== 'admin') {
      return <Navigate to="/login" />;
    }
  const [formData, setFormData] = useState<Partial<Job>>({
    title: '',
    company: '',
    location: '',
    type: '',
    salary: '',
    description: '',
    requirements: [],
    responsibilities: [],
    benefits: [],
    category: '',
  });

  useEffect(() => {
    const loadJob = async () => {
      if (!id) return;
      try {
        const jobData = await fetchJobById(id);
        setJob(jobData);
        setFormData({
          title: jobData.title,
          company: jobData.company,
          location: jobData.location,
          type: jobData.type,
          salary: jobData.salary,
          description: jobData.description,
          requirements: jobData.requirements,
          responsibilities: jobData.responsibilities,
          benefits: jobData.benefits,
          category: jobData.category,
        });
      } catch (err) {
        setError('Failed to load job details. Please try again.');
      } finally {
        setLoading(false);
      }
    };
    loadJob();
  }, [id]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleArrayInputChange = (e: React.ChangeEvent<HTMLInputElement>, field: string) => {
    const value = e.target.value.split(',').map((item) => item.trim());
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!id) return;
    try {
      await updateJob(id, formData);
      navigate('/admin/dashboard');
    } catch (err) {
      setError('Failed to update job. Please try again.');
    }
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
  if (!job) return <div>Job not found</div>;

  return (
    <div className="max-w-3xl mx-auto p-6 bg-white rounded-lg shadow-md">
      <h2 className="text-2xl font-bold mb-4">Edit Job</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium">Title</label>
          <input
            type="text"
            name="title"
            value={formData.title || ''}
            onChange={handleInputChange}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Company</label>
          <input
            type="text"
            name="company"
            value={formData.company || ''}
            onChange={handleInputChange}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Location</label>
          <input
            type="text"
            name="location"
            value={formData.location || ''}
            onChange={handleInputChange}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Type</label>
          <input
            type="text"
            name="type"
            value={formData.type || ''}
            onChange={handleInputChange}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Salary</label>
          <input
            type="text"
            name="salary"
            value={formData.salary || ''}
            onChange={handleInputChange}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Description</label>
          <textarea
            name="description"
            value={formData.description || ''}
            onChange={handleInputChange}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Requirements (comma-separated)</label>
          <input
            type="text"
            value={formData.requirements?.join(', ') || ''}
            onChange={(e) => handleArrayInputChange(e, 'requirements')}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Responsibilities (comma-separated)</label>
          <input
            type="text"
            value={formData.responsibilities?.join(', ') || ''}
            onChange={(e) => handleArrayInputChange(e, 'responsibilities')}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Benefits (comma-separated)</label>
          <input
            type="text"
            value={formData.benefits?.join(', ') || ''}
            onChange={(e) => handleArrayInputChange(e, 'benefits')}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Category</label>
          <input
            type="text"
            name="category"
            value={formData.category || ''}
            onChange={handleInputChange}
            className="mt-1 block w-full border-gray-300 rounded-md p-2"
          />
        </div>
        <button type="submit" className="w-full py-2 px-4 bg-blue-600 text-white rounded-md">
          Save Changes
        </button>
      </form>
    </div>
  );
};

export default EditJobPage;