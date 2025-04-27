import React, { useState, useEffect } from 'react';
import { Plus, Edit, Trash2, BarChart2, FileText, Users } from 'lucide-react';
import { Link } from 'react-router-dom';
import { formatDate } from '../../../frontend/utils/formatters';
import { getJobStatusBadgeClass } from '../../../frontend/utils/statusHelpers';
import Button from '../../../frontend/components/ui/Button';
import Card, { CardHeader, CardBody } from '../../../frontend/components/ui/Card';
import JobFilter from '../../../frontend/components/jobs/JobFilter';
import Badge from '../../../frontend/components/ui/Badge';
import { fetchAdminJobs, deleteJob, fetchUserApplications } from '../../../frontend/services/api';
import { useAuth } from '../../../frontend/context/AuthContext';

const AdminDashboardPage: React.FC = () => {
  const { user } = useAuth();
  const [activeTab, setActiveTab] = useState<'overview' | 'jobs' | 'applications'>('overview');
  const [searchTerm, setSearchTerm] = useState('');
  const [filterOpen, setFilterOpen] = useState(false);
  const [selectedStatus, setSelectedStatus] = useState('all');
  const [jobs, setJobs] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [selectedJobId, setSelectedJobId] = useState<string | null>(null);

  useEffect(() => {
    const fetchJobs = async () => {
      try {
        setLoading(true);
        setError(null);
        
        const jobsData = await fetchAdminJobs();
        setJobs(jobsData);
        
        setLoading(false);
      } catch (err) {
        console.error('Error fetching admin jobs:', err);
        setError('Failed to load jobs. Please try again later.');
        setLoading(false);
      }
    };
    
    fetchJobs();
  }, []);

  // Filter jobs based on search term and status
  // const filteredJobs = jobs.filter(job => {
  //   const matchesSearch = job.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
  //                        job.company.toLowerCase().includes(searchTerm.toLowerCase()) ||
  //                        job.location.toLowerCase().includes(searchTerm.toLowerCase());
    
  //   const matchesStatus = selectedStatus === 'all' || job.status === selectedStatus;
    
  //   return matchesSearch && matchesStatus;
  // });

  // Calculate statistics
  const totalJobs = jobs.length;
  const activeJobs = jobs.filter(job => job.status === 'active').length;
  const totalApplications = jobs.reduce((sum, job) => sum + job.applications, 0);
  const [applications, setApplications] = useState<any[]>([]);
  
  // Group jobs by category for chart
  const jobsByCategory = jobs.reduce((acc, job) => {
    acc[job.category] = (acc[job.category] || 0) + 1;
    return acc;
  }, {} as Record<string, number>);

  const handleDeleteJob = (jobId: string) => {
    setSelectedJobId(jobId);
    setShowDeleteConfirm(true);
  };

  console.log('application', applications);

  useEffect(() => {
    const fetchApplications = async () => {
      try {
        setLoading(true);
        setError(null);
        
        const applicationsData = await fetchUserApplications(); // Fetching user applications
        setApplications(applicationsData);
        
        setLoading(false);
      } catch (err) {
        console.error('Error fetching applications:', err);
        setError('Failed to load applications. Please try again later.');
        setLoading(false);
      }
    };
    fetchApplications();
    
    // if (activeTab === 'applications') {
    //   fetchApplications();
    // }
  }, []);

  const confirmDeleteJob = async () => {
    if (!selectedJobId) return;
    
    try {
      await deleteJob(selectedJobId);
      
      // Remove job from state
      setJobs(jobs.filter(job => job.id !== selectedJobId));
      
      setShowDeleteConfirm(false);
      setSelectedJobId(null);
    } catch (err) {
      console.error('Error deleting job:', err);
      // Show error message
    }
  };

  // Get category color for chart
  const getCategoryColor = (category: string): string => {
    switch (category) {
      case 'Technology':
        return 'bg-blue-500';
      case 'Healthcare':
        return 'bg-green-500';
      case 'Education':
        return 'bg-orange-500';
      case 'Business':
        return 'bg-purple-500';
      case 'Creative':
        return 'bg-pink-500';
      case 'Hospitality':
        return 'bg-yellow-500';
      default:
        return 'bg-gray-500';
    }
  };

  return (
    <div className="bg-gray-50 min-h-screen py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold">Admin Dashboard</h1>
          <Button 
            variant="secondary" 
            icon={<Plus className="h-5 w-5" />}
            href="/admin/jobs/new"
          >
            Post New Job
          </Button>
        </div>
        
        {/* Dashboard Tabs */}
        <div className="bg-white rounded-lg shadow-md mb-8">
          <div className="border-b border-gray-200">
            <nav className="flex -mb-px">
              <button
                className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
                  activeTab === 'overview'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
                onClick={() => setActiveTab('overview')}
              >
                <BarChart2 className="h-5 w-5 inline-block mr-2" />
                Overview
              </button>
              <button
                className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
                  activeTab === 'jobs'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
                onClick={() => setActiveTab('jobs')}
              >
                <FileText className="h-5 w-5 inline-block mr-2" />
                Jobs Management
              </button>
              <button
                className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
                  activeTab === 'applications'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
                onClick={() => setActiveTab('applications')}
              >
                <Users className="h-5 w-5 inline-block mr-2" />
                Applications
              </button>
            </nav>
          </div>
          
          <div className="p-6">
            {loading ? (
              <div className="text-center py-12">
                <div className="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-blue-600 border-r-transparent"></div>
                <p className="mt-4 text-gray-600">Loading dashboard data...</p>
              </div>
            ) : error ? (
              <div className="text-center py-12 bg-red-50 rounded-lg">
                <p className="text-red-600 mb-4">{error}</p>
                <Button 
                  onClick={() => window.location.reload()}
                  variant="primary"
                >
                  Try Again
                </Button>
              </div>
            ) : activeTab === 'overview' ? (
              <div>
                {/* Stats Cards */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
                  <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                    <h3 className="text-lg font-medium text-gray-500">Total Jobs</h3>
                    <p className="text-3xl font-bold">{totalJobs}</p>
                  </div>
                  <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                    <h3 className="text-lg font-medium text-gray-500">Active Jobs</h3>
                    <p className="text-3xl font-bold">{activeJobs}</p>
                  </div>
                  <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                    <h3 className="text-lg font-medium text-gray-500">Total Applications</h3>
                    <p className="text-3xl font-bold">{totalApplications}</p>
                  </div>
                  <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                    <h3 className="text-lg font-medium text-gray-500">Avg. Applications</h3>
                    <p className="text-3xl font-bold">{totalJobs > 0 ? Math.round(totalApplications / totalJobs) : 0}</p>
                  </div>
                </div>
                
                {/* Jobs by Category */}
                <div className="bg-white p-6 rounded-lg shadow border border-gray-200 mb-8">
                  <h3 className="text-lg font-medium mb-4">Jobs by Category</h3>
                  <div className="space-y-4">
                    {Object.entries(jobsByCategory).map(([category, count]) => (
                      <div key={category} className="flex items-center">
                        <div className="w-32 text-sm">{category}</div>
                        <div className="flex-1 h-5 bg-gray-200 rounded-full overflow-hidden">
                          <div 
                            className={`h-full ${getCategoryColor(category)}`} 
                            style={{ width: `${(count / totalJobs) * 100}%` }}
                          ></div>
                        </div>
                        <div className="w-10 text-right text-sm">{count}</div>
                      </div>
                    ))}
                  </div>
                </div>
              </div>
            ) : activeTab === 'jobs' ? (
              <div>
                {/* Search and Filter */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
          {loading ? (
            <div className="text-center py-12">
              <div className="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-blue-600 border-r-transparent"></div>
              <p className="mt-4 text-gray-600">Loading jobs...</p>
            </div>
          ) : error ? (
            <div className="text-center py-12 bg-red-50 rounded-lg">
              <p className="text-red-600 mb-4">{error}</p>
              <button
                onClick={() => window.location.reload()}
                className="px-4 py-2 bg-blue-600 text-white rounded-md"
              >
                Try Again
              </button>
            </div>
          ) : (
            jobs.map((job) => (
              <Card key={job.id} className="bg-white p-6 rounded-lg shadow-lg">
                <CardHeader>
                  <h3 className="text-xl font-semibold text-gray-900">{job.title}</h3>
                  <p className="text-sm text-gray-600">{job.company} - {job.location}</p>
                </CardHeader>
                <CardBody>
                  <p className="text-md text-gray-700 mt-2">{job.description}</p>
                  <p className="text-sm text-gray-600 mt-4">Salary: {job.salary}</p>
                  <a
                    href={`/admin/jobs/${job.id}`}
                    className="text-blue-500 mt-4 inline-block"
                  >
                    View Job Details
                  </a>
                  <button
                    onClick={() => handleDeleteJob(job.id)}
                    className="ml-4 text-red-500 hover:text-red-700"
                  >
                    <Trash2 className="h-5 w-5 inline-block" />
                    Delete Job
                  </button>
                </CardBody>
              </Card>
            ))
          )}
        </div>
                {/* <JobFilter 
                  searchTerm={searchTerm}
                  onSearchChange={(e) => setSearchTerm(e.target.value)}
                  isFilterOpen={filterOpen}
                  toggleFilter={() => setFilterOpen(!filterOpen)}
                  selectedStatus={selectedStatus}
                  onStatusChange={(e) => setSelectedStatus(e.target.value)}
                /> */}
              </div>
            ) : (
              <div>
                {/* Display Applications in Cards */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
                  {applications.map((application) => (
                    <Card key={application.id} className="bg-white p-6 rounded-lg shadow-lg">
                      <CardHeader>
                        <h3 className="text-xl font-semibold text-gray-900">{application.full_name}</h3>
                        <p className="text-sm text-gray-600">{application.email}</p>
                      </CardHeader>
                      <CardBody>
                        <p className="text-md text-gray-700 mt-2">Applied for: {application.jobTitle}</p>
                        <p className="text-sm text-gray-600 mt-4">Status: {application.status}</p>
                        <a
                          href={`/admin/applications/${application.id}`}
                          className="text-blue-500 mt-4 inline-block"
                        >
                          View Application Details
                        </a>
                      </CardBody>
                    </Card>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default AdminDashboardPage;