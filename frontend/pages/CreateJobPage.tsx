import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createJob } from '../../../frontend/services/api';

import { useAuth } from '../../../frontend/context/AuthContext';
import { Navigate } from 'react-router-dom';

const EditJobPage = () => {
  const { isAuthenticated, user } = useAuth();

  if (!isAuthenticated || user?.role !== 'admin') {
    return <Navigate to="/login" />;
  }
};

const CreateJobPage: React.FC = () => {

  
  const navigate = useNavigate();

  // Form State
  const [form, setForm] = useState({
    title: '',
    company: '',
    location: '',
    type: 'Full-time',
    salary: '',
    description: '',
    category: 'Technology',
    companyInfo: {
      name: '',
      description: '',
      website: '',
      industry: '',
      size: ''
    },
    requirements: '',
    responsibilities: '',
    benefits: ''
  });
  
enum JobCategory {
  Technology = "Technology",
  Healthcare = "Healthcare",
  Education = "Education",
  Business = "Business",
  Creative = "Creative",
  Hospitality = "Hospitality"
};

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);

  // Handle Form Change
  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    if (name.startsWith('companyInfo.')) {
      setForm({
        ...form,
        companyInfo: {
          ...form.companyInfo,
          [name.split('.')[1]]: value
        }
      });
    } else {
      setForm({ ...form, [name]: value });
    }
  };

  // Handle Submit
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
  
    if (!form.title || !form.company || !form.location || !form.description) {
      setError('Please fill in all required fields.');
      return;
    }
  
    setLoading(true);
    setError(null);
  
    // Ensure category is properly typed
    const formattedJob = {
      ...form,
      category: form.category as JobCategory, // ✅ Fix category type
      requirements: form.requirements.split(',').map(item => item.trim()).filter(Boolean), // ✅ Convert to array
      responsibilities: form.responsibilities.split(',').map(item => item.trim()).filter(Boolean), // ✅ Convert to array
      benefits: form.benefits.split(',').map(item => item.trim()).filter(Boolean), // ✅ Convert to array
    };
  
    console.log('🚀 Job Data:', JSON.stringify(formattedJob, null, 2));
  
    try {
      await createJob(formattedJob);
      setSuccess(true);
      setTimeout(() => navigate('/admin/dashboard'), 1500);
    } catch (err) {
      console.error('❌ Job creation error:', err);
      setError('Failed to create job. Please try again.');
    } finally {
      setLoading(false);
    }
  };
  

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="max-w-2xl w-full bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-2xl font-bold text-center mb-4">Create Job</h2>

        {error && <p className="text-red-600 text-center mb-4">{error}</p>}
        {success && <p className="text-green-600 text-center mb-4">Job created successfully! Redirecting...</p>}

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Job Title */}
          <Input label="Job Title" name="title" value={form.title} onChange={handleChange} required />

          {/* Company */}
          <Input label="Company" name="company" value={form.company} onChange={handleChange} required />

          {/* Location */}
          <Input label="Location" name="location" value={form.location} onChange={handleChange} required />

          {/* Type & Category */}
          <div className="grid grid-cols-2 gap-4">
            <Select label="Job Type" name="type" options={['Full-time', 'Part-time', 'Contract', 'Internship']} value={form.type} onChange={handleChange} />
            <Select label="Category" name="category" options={['Technology', 'Healthcare', 'Education', 'Business', 'Creative', 'Hospitality']} value={form.category} onChange={handleChange} />
          </div>

          {/* Salary */}
          <Input label="Salary" name="salary" value={form.salary} onChange={handleChange} />

          {/* Description */}
          <TextArea label="Job Description" name="description" value={form.description} onChange={handleChange} required />

          {/* Requirements, Responsibilities, Benefits (Comma-separated) */}
          <TextArea label="Requirements (comma-separated)" name="requirements" value={form.requirements} onChange={handleChange} />
          <TextArea label="Responsibilities (comma-separated)" name="responsibilities" value={form.responsibilities} onChange={handleChange} />
          <TextArea label="Benefits (comma-separated)" name="benefits" value={form.benefits} onChange={handleChange} />

          <h3 className="text-lg font-semibold mt-4">Company Information</h3>

          <Input label="Company Name" name="companyInfo.name" value={form.companyInfo.name} onChange={handleChange} required />
          <TextArea label="Company Description" name="companyInfo.description" value={form.companyInfo.description} onChange={handleChange} required />
          <Input label="Website" name="companyInfo.website" value={form.companyInfo.website} onChange={handleChange} />
          <Input label="Industry" name="companyInfo.industry" value={form.companyInfo.industry} onChange={handleChange} />
          <Input label="Company Size" name="companyInfo.size" value={form.companyInfo.size} onChange={handleChange} />

          {/* Submit Button */}
          <button type="submit" className="w-full py-2 px-4 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700">
            {loading ? 'Creating...' : 'Create Job'}
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateJobPage;

// Reusable Input Component
const Input: React.FC<{ label: string; name: string; value: string; onChange: (e: React.ChangeEvent<HTMLInputElement>) => void; required?: boolean }> = ({ label, name, value, onChange, required }) => (
  <div>
    <label className="block text-sm font-medium text-gray-700">{label}</label>
    <input type="text" name={name} value={value} onChange={onChange} required={required} className="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2" />
  </div>
);

// Reusable TextArea Component
const TextArea: React.FC<{ label: string; name: string; value: string; onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void; required?: boolean }> = ({ label, name, value, onChange, required }) => (
  <div>
    <label className="block text-sm font-medium text-gray-700">{label}</label>
    <textarea name={name} value={value} onChange={onChange} required={required} className="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2"></textarea>
  </div>
);

// Reusable Select Input Component
const Select: React.FC<{ label: string; name: string; options: string[]; value: string; onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void }> = ({ label, name, options, value, onChange }) => (
  <div>
    <label className="block text-sm font-medium text-gray-700">{label}</label>
    <select name={name} value={value} onChange={onChange} className="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2">
      {options.map((option) => (
        <option key={option} value={option}>{option}</option>
      ))}
    </select>
  </div>
);
