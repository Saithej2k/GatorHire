import React from 'react';
import { Filter, ChevronDown } from 'lucide-react';
import SearchInput from '../ui/SearchInput';

interface JobFilterProps {
  searchTerm: string;
  onSearchChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  isFilterOpen: boolean;
  toggleFilter: () => void;
  selectedJobType?: string;
  onJobTypeChange?: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  selectedLocation?: string;
  onLocationChange?: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  selectedStatus?: string;
  onStatusChange?: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  showStatusFilter?: boolean;
}

const JobFilter: React.FC<JobFilterProps> = ({
  searchTerm,
  onSearchChange,
  isFilterOpen,
  toggleFilter,
  selectedJobType,
  onJobTypeChange,
  selectedLocation,
  onLocationChange,
  selectedStatus,
  onStatusChange,
  showStatusFilter = false,
}) => {
  return (
    <div className="bg-white p-4 rounded-lg shadow-md mb-8">
      <div className="flex flex-col md:flex-row gap-4 mb-4">
        <div className="flex-grow">
          <SearchInput
            value={searchTerm}
            onChange={onSearchChange}
            placeholder="Search jobs, companies, or locations"
          />
        </div>
        <button 
          className="md:w-auto flex items-center justify-center gap-2 bg-gray-100 hover:bg-gray-200 px-4 py-2 rounded-md"
          onClick={toggleFilter}
        >
          <Filter className="h-5 w-5" />
          <span>Filters</span>
        </button>
      </div>
      
      {isFilterOpen && (
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mt-4 pt-4 border-t border-gray-200">
          {showStatusFilter && (
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Job Status</label>
              <div className="relative">
                <select
                  className="w-full border border-gray-300 rounded-md p-2 pr-8 focus:outline-none focus:ring-2 focus:ring-blue-500 appearance-none"
                  value={selectedStatus}
                  onChange={onStatusChange}
                >
                  <option value="all">All Statuses</option>
                  <option value="active">Active</option>
                   <option value="closed">Closed</option>
                  <option value="draft">Draft</option>
                </select>
                <ChevronDown className="absolute right-2 top-1/2 transform -translate-y-1/2 text-gray-400 h-5 w-5 pointer-events-none" />
              </div>
            </div>
          )}
          
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Job Type</label>
            <div className="relative">
              <select
                className="w-full border border-gray-300 rounded-md p-2 pr-8 focus:outline-none focus:ring-2 focus:ring-blue-500 appearance-none"
                value={selectedJobType}
                onChange={onJobTypeChange}
              >
                <option value="">All Types</option>
                <option value="full-time">Full-time</option>
                <option value="part-time">Part-time</option>
                <option value="contract">Contract</option>
                <option value="internship">Internship</option>
              </select>
              <ChevronDown className="absolute right-2 top-1/2 transform -translate-y-1/2 text-gray-400 h-5 w-5 pointer-events-none" />
            </div>
          </div>
          
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Location</label>
            <div className="relative">
              <select
                className="w-full border border-gray-300 rounded-md p-2 pr-8 focus:outline-none focus:ring-2 focus:ring-blue-500 appearance-none"
                value={selectedLocation}
                onChange={onLocationChange}
              >
                <option value="">All Locations</option>
                <option value="remote">Remote</option>
                <option value="san-francisco">San Francisco, CA</option>
                <option value="new-york">New York, NY</option>
                <option value="seattle">Seattle, WA</option>
                <option value="austin">Austin, TX</option>
              </select>
              <ChevronDown className="absolute right-2 top-1/2 transform -translate-y-1/2 text-gray-400 h-5 w-5 pointer-events-none" />
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default JobFilter;