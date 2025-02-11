import React from 'react';
import { Search, BookmarkPlus, Bell, Settings } from 'lucide-react';

export const Dashboard = () => {
  const recentJobs = [
    {
      id: 1,
      title: 'Senior Software Engineer',
      company: 'Tech Corp',
      location: 'San Francisco, CA',
      salary: '$120k - $160k',
      type: 'Full-time',
    },
    {
      id: 2,
      title: 'Product Manager',
      company: 'Innovation Labs',
      location: 'New York, NY',
      salary: '$100k - $130k',
      type: 'Full-time',
    },
    {
      id: 3,
      title: 'UX Designer',
      company: 'Design Studio',
      location: 'Remote',
      salary: '$90k - $120k',
      type: 'Contract',
    },
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Left Sidebar */}
          <div className="lg:col-span-1">
            <div className="bg-white rounded-lg shadow-md p-6">
              <div className="flex items-center space-x-4 mb-6">
                <div className="bg-orange-500 rounded-full w-12 h-12 flex items-center justify-center text-white text-xl font-bold">
                  JD
                </div>
                <div>
                  <h2 className="text-lg font-semibold text-blue-900">John Doe</h2>
                  <p className="text-gray-600">Software Engineer</p>
                </div>
              </div>
              <nav className="space-y-2">
                <a href="#" className="flex items-center space-x-2 text-gray-600 hover:text-orange-500 p-2 rounded-md hover:bg-orange-50">
                  <Search className="h-5 w-5" />
                  <span>Job Search</span>
                </a>
                <a href="#" className="flex items-center space-x-2 text-gray-600 hover:text-orange-500 p-2 rounded-md hover:bg-orange-50">
                  <BookmarkPlus className="h-5 w-5" />
                  <span>Saved Jobs</span>
                </a>
                <a href="#" className="flex items-center space-x-2 text-gray-600 hover:text-orange-500 p-2 rounded-md hover:bg-orange-50">
                  <Bell className="h-5 w-5" />
                  <span>Job Alerts</span>
                </a>
                <a href="#" className="flex items-center space-x-2 text-gray-600 hover:text-orange-500 p-2 rounded-md hover:bg-orange-50">
                  <Settings className="h-5 w-5" />
                  <span>Settings</span>
                </a>
              </nav>
            </div>
          </div>

          {/* Main Content */}
          <div className="lg:col-span-2">
            <div className="bg-white rounded-lg shadow-md p-6 mb-8">
              <h2 className="text-xl font-semibold text-blue-900 mb-4">Recent Job Listings</h2>
              <div className="space-y-6">
                {recentJobs.map((job) => (
                  <div key={job.id} className="border-b pb-6 last:border-b-0 last:pb-0">
                    <div className="flex justify-between items-start">
                      <div>
                        <h3 className="text-lg font-medium text-blue-900">{job.title}</h3>
                        <p className="text-gray-600">{job.company}</p>
                        <div className="mt-2 flex items-center space-x-4">
                          <span className="text-sm text-gray-500">{job.location}</span>
                          <span className="text-sm text-gray-500">{job.salary}</span>
                          <span className="text-sm text-gray-500">{job.type}</span>
                        </div>
                      </div>
                      <button className="bg-orange-500 text-white px-4 py-2 rounded-md hover:bg-orange-600">
                        Apply Now
                      </button>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};