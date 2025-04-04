import React from 'react';
import { Link } from 'react-router-dom';
import { MapPin, Briefcase, Clock } from 'lucide-react';
import { Job } from '../../types/job';
import { formatDate } from '../../utils/formatters';
import Button from '../ui/Button';
import Badge from '../ui/Badge';

interface JobCardProps {
  job: Job;
}

const JobCard: React.FC<JobCardProps> = ({ job }) => {
  // Get category badge color
  const getCategoryBadgeClass = (category: string) => {
    switch (category) {
      case 'Technology':
        return 'bg-blue-100 text-blue-800';
      case 'Healthcare':
        return 'bg-green-100 text-green-800';
      case 'Education':
        return 'bg-orange-100 text-orange-800';
      case 'Business':
        return 'bg-purple-100 text-purple-800';
      case 'Creative':
        return 'bg-pink-100 text-pink-800';
      case 'Hospitality':
        return 'bg-yellow-100 text-yellow-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <div data-testid="job-listing" className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition duration-300">
      <div className="p-6">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-4">
          <h2 className="text-xl font-bold text-gray-900 mb-1 md:mb-0">{job.title}</h2>
          <div className="flex flex-wrap gap-2">
            <span className="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-blue-100 text-blue-800">
              {job.type}
            </span>
            <Badge 
              text={job.category}
              className={getCategoryBadgeClass(job.category)}
            />
          </div>
        </div>
        
        <div className="mb-4">
          <div className="flex items-center text-gray-600 mb-2">
            <Briefcase className="h-4 w-4 mr-2" />
            <span>{job.company}</span>
          </div>
          <div className="flex items-center text-gray-600 mb-2">
            <MapPin className="h-4 w-4 mr-2" />
            <span>{job.location}</span>
          </div>
          <div className="flex items-center text-gray-600">
            <Clock className="h-4 w-4 mr-2" />
            <span>Posted on {formatDate(job.postedDate)}</span>
          </div>
        </div>
        
        <div className="mb-4">
          <p className="text-gray-700">Job ID: {job.id}</p>
        </div>

        <div className="mb-4">
          <p className="text-gray-700">{job.description.substring(0, 150)}...</p>
        </div>
        
        <div className="flex flex-wrap gap-2 mb-4">
          {job.requirements.slice(0, 3).map((req, index) => (
            <span key={index} className="bg-gray-100 text-gray-800 text-xs px-2 py-1 rounded">
              {req}
            </span>
          ))}
        </div>
        
        <div className="flex items-center justify-between">
          <span className="text-gray-900 font-medium">{job.salary}</span>
          <Button href={`/jobs/${job.id}`} variant="primary" size="sm">
            View Details
          </Button>
        </div>
      </div>
    </div>
  );
};

export default JobCard;