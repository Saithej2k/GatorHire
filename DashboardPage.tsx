import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Briefcase, MapPin, Clock, BookOpen, AlertCircle } from 'lucide-react';
import { useAuth } from '../frontend/context/AuthContext';
import { formatDate } from '../frontend/utils/formatters';
import { getApplicationStatusBadgeClass } from '../frontend/utils/statusHelpers';
import Card, { CardHeader, CardBody } from '../frontend/components/ui/Card';
import Badge from '../frontend/components/ui/Badge';
import Button from '../frontend/components/ui/Button';
import { Application } from '../frontend/types/application';
import { SavedJob } from '../frontend/types/job';
import { fetchUserApplications, fetchSavedJobs, fetchJobs } from '../frontend/services/api';

const DashboardPage: React.FC = () => {
  const { user } = useAuth();
  const [activeTab, setActiveTab] = useState<'applications' | 'saved'>('applications');
  const [applications, setApplications] = useState<Application[]>([]);
  const [savedJobs, setSavedJobs] = useState<SavedJob[]>([]);
  const [recommendedJobs, setRecommendedJobs] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [usingMockData, setUsingMockData] = useState<boolean>(false);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        setLoading(true);
        setError(null);
        setUsingMockData(false);
        
        // Fetch user applications
        const userApplications = await fetchUserApplications();
        console.log("User Applications:", userApplications); // Add logging
        setApplications(userApplications);
        
        // Fetch saved jobs
        const userSavedJobs = await fetchSavedJobs();
        setSavedJobs(userSavedJobs);
        
        // // Fetch recommended jobs based on user's profile
        // // In a real app, this would use a recommendation algorithm
        // // For now, just fetch some jobs that aren't saved or applied to
        const allJobs = await fetchJobs();
        
        // // Check if we're using mock data
        if (userApplications.length === 0 && userSavedJobs.length === 0 && allJobs.length > 0) {
          setUsingMockData(true);
        }
        
        const appliedJobIds = new Set(userApplications.map(app => app.jobId));
        const savedJobIds = new Set(userSavedJobs.map(job => job.id));
        
        // Filter jobs that user hasn't interacted with
        // const filteredJobs = allJobs.filter(job => 
        //   !appliedJobIds.has(job.id) && !savedJobIds.has(job.id)
        // ).slice(0, 3); // Just take the first 3
        const filteredJobs = allJobs.filter(job => 
          !appliedJobIds.has(job.id)
        ).slice(0, 3); // Just take the first 3
        
        setRecommendedJobs(filteredJobs);
        setLoading(false);
      } catch (err) {
        console.error('Error fetching user data:', err);
        setError('Failed to load your data. Using sample data instead.');
        setUsingMockData(true);
        setLoading(false);
      }
    };
    
    fetchUserData();
  }, []);

  return (
    <div className="bg-gray-50 min-h-screen py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="mb-8">
          <h1 className="text-3xl font-bold">Dashboard</h1>
          <p className="text-gray-600 mt-2">Welcome back, {user?.fullName || 'User'}!</p>
        </div>
        
        {usingMockData && (
          <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4 mb-6">
            <p className="text-yellow-700 text-sm">
              Note: Showing sample data. The server is currently unavailable.
            </p>
          </div>
        )}
        
        {/* Dashboard Tabs */}
        <Card className="mb-8">
          <div className="border-b border-gray-200">
            <nav className="flex -mb-px">
              <button
                className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
                  activeTab === 'applications'
                    ? 'border-orange-500 text-orange-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
                onClick={() => setActiveTab('applications')}
              >
                <Briefcase className="h-5 w-5 inline-block mr-2" />
                My Applications
              </button>
              <button
                className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
                  activeTab === 'saved'
                    ? 'border-orange-500 text-orange-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
                onClick={() => setActiveTab('saved')}
              >
                <BookOpen className="h-5 w-5 inline-block mr-2" />
                Saved Jobs
              </button>
            </nav>
          </div>
          
          <CardBody>
            {loading ? (
              <div className="text-center py-12">
                <div className="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-orange-600 border-r-transparent"></div>
                <p className="mt-4 text-gray-600">Loading your data...</p>
              </div>
            ) : error ? (
              <div className="text-center py-6 bg-yellow-50 rounded-lg">
                <p className="text-yellow-700 mb-4">{error}</p>
              </div>
            ) : activeTab === 'applications' ? (
              <>
                <h2 className="text-xl font-semibold mb-4">Your Job Applications</h2>
                {applications.length === 0 ? (
                  <div className="text-center py-8 bg-gray-50 rounded-lg">
                    <AlertCircle className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                    <p className="text-gray-600 mb-4">You haven't applied to any jobs yet.</p>
                    <Button href="/jobs" variant="primary">
                      Browse Jobs
                    </Button>
                  </div>
                ) : (
                  <div className="overflow-x-auto">
                    <table className="min-w-full divide-y divide-gray-200">
                      <thead className="bg-gray-50">
                        <tr>
                          <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Job
                          </th>
                          <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Applied Date
                          </th>
                          <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Status
                          </th>
                          <th scope="col" className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Actions
                          </th>
                        </tr>
                      </thead>
                      <tbody className="bg-white divide-y divide-gray-200">
                        {applications.map((application) => (
                          <tr key={application.id} className="hover:bg-gray-50">
                            <td className="px-6 py-4 whitespace-nowrap">
                              <div className="flex items-center">
                                <div>
                                  <div className="text-sm font-medium text-gray-900">{application.jobTitle}</div>
                                  <div className="text-sm text-gray-500">{application.company}</div>
                                </div>
                              </div>
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap">
                              <div className="text-sm text-gray-900">{formatDate(application.appliedDate)}</div>
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap">
                              <Badge 
                                text={application.status.charAt(0).toUpperCase() + application.status.slice(1)}
                                className={getApplicationStatusBadgeClass(application.status)}
                              />
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                              <Link to={`/jobs/${application.jobId}`} className="text-blue-600 hover:text-blue-900">
                                View Job
                              </Link>
                            </td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                )}
              </>
            ) : (
              <>
                <h2 className="text-xl font-semibold mb-4">Your Saved Jobs</h2>
                {savedJobs.length === 0 ? (
                  <div className="text-center py-8 bg-gray-50 rounded-lg">
                    <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                    <p className="text-gray-600 mb-4">You haven't saved any jobs yet.</p>
                    <Button href="/jobs" variant="primary">
                      Browse Jobs
                    </Button>
                  </div>
                ) : (
                  <div className="space-y-4">
                    {savedJobs.map((job) => (
                      <div key={job.id} className="border border-gray-200 rounded-lg p-4 hover:shadow-md transition duration-200">
                        <div className="flex flex-col md:flex-row md:items-center md:justify-between">
                          <div>
                            <h3 className="text-lg font-medium text-gray-900">{job.title}</h3>
                            <div className="mt-1 flex flex-wrap gap-y-1 gap-x-4 text-sm text-gray-500">
                              <div className="flex items-center">
                                <Briefcase className="h-4 w-4 mr-1" />
                                {job.company}
                              </div>
                              <div className="flex items-center">
                                <MapPin className="h-4 w-4 mr-1" />
                                {job.location}
                              </div>
                              <div className="flex items-center">
                                <Clock className="h-4 w-4 mr-1" />
                                Saved on {formatDate(job.savedDate)}
                              </div>
                            </div>
                          </div>
                          <div className="mt-4 md:mt-0 flex flex-col sm:flex-row gap-2">
                            <Button href={`/jobs/${job.id}`} variant="primary" size="sm">
                              View Details
                            </Button>
                            <Button href={`/apply/${job.id}`} variant="outline" size="sm">
                              Apply Now
                            </Button>
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                )}
              </>
            )}
          </CardBody>
        </Card>
        
        {/* Recommended Jobs */}
        <Card>
          <CardHeader>
            <h2 className="text-xl font-semibold">Recommended Jobs</h2>
            <p className="text-gray-600 mt-1">Based on your profile and preferences</p>
          </CardHeader>
          
          <CardBody>
            {loading ? (
              <div className="text-center py-6">
                <div className="inline-block h-6 w-6 animate-spin rounded-full border-4 border-solid border-orange-600 border-r-transparent"></div>
                <p className="mt-2 text-gray-600">Loading recommendations...</p>
              </div>
            ) : recommendedJobs.length === 0 ? (
              <div className="text-center py-6">
                <p className="text-gray-600">No recommendations available at this time.</p>
              </div>
            ) : (
              <div className="space-y-6">
                {recommendedJobs.map(job => (
                  <div key={job.id} className="border border-gray-200 rounded-lg p-4 hover:shadow-md transition duration-200">
                    <div className="flex flex-col md:flex-row md:items-center md:justify-between">
                      <div>
                        <h3 className="text-lg font-medium text-gray-900">{job.title}</h3>
                        <div className="mt-1 flex flex-wrap gap-y-1 gap-x-4 text-sm text-gray-500">
                          <div className="flex items-center">
                            <Briefcase className="h-4 w-4 mr-1" />
                            {job.company}
                          </div>
                          <div className="flex items-center">
                            <MapPin className="h-4 w-4 mr-1" />
                            {job.location}
                          </div>
                          <div className="flex items-center">
                            <Clock className="h-4 w-4 mr-1" />
                            Posted {formatDate(job.postedDate)}
                          </div>
                        </div>
                      </div>
                      <div className="mt-4 md:mt-0">
                        <Button href={`/jobs/${job.id}`} variant="primary" size="sm">
                          View Details
                        </Button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </CardBody>
        </Card>
      </div>
    </div>
  );
};

export default DashboardPage;