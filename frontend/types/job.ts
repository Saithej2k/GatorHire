/**
 * Types related to job listings
 */

export interface Job {
  id: string;
  title: string;
  company: string;
  location: string;
  type: string;
  salary: string;
  description: string;
  requirements: string[];
  responsibilities?: string[];
  benefits?: string[];
  postedDate: string;
  companyInfo?: CompanyInfo;
  category: JobCategory;
  status?: string;
}

export type JobCategory = 'Technology' | 'Healthcare' | 'Education' | 'Business' | 'Creative' | 'Hospitality' | 'All';

export interface CompanyInfo {
  name: string;
  description: string;
  website?: string;
  industry?: string;
  size?: string;
}

export interface SavedJob {
  id: string;
  title: string;
  company: string;
  location: string;
  type: string;
  postedDate: string;
  savedDate: string;
  category: JobCategory;
}

export interface AdminJob extends Job {
  status: 'active' | 'closed' | 'draft';
  applications: number;
}