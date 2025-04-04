import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import { describe, it, expect } from 'vitest';
import JobCard from '../components/jobs/JobCard';
import { Job } from '../types/job';

// Mock job data
const mockJob: Job = {
  id: '1',
  title: 'Senior Frontend Developer',
  company: 'TechCorp',
  location: 'San Francisco, CA',
  type: 'Full-time',
  salary: '$120,000 - $150,000',
  description: 'We are looking for an experienced Frontend Developer to join our team...',
  requirements: ['React', 'TypeScript', 'Redux'],
  postedDate: '2023-04-15',
  category: 'Technology'
};

// Mock router
const renderWithRouter = (ui: React.ReactElement) => {
  return render(ui, { wrapper: BrowserRouter });
};

describe('JobCard Component', () => {
  it('renders job information correctly', () => {
    renderWithRouter(<JobCard job={mockJob} />);
    
    // Check if job title is displayed
    expect(screen.getByText(mockJob.title)).toBeInTheDocument();
    
    // Check if company name is displayed
    expect(screen.getByText(mockJob.company)).toBeInTheDocument();
    
    // Check if location is displayed
    expect(screen.getByText(mockJob.location)).toBeInTheDocument();
    
    // Check if job type is displayed
    expect(screen.getByText(mockJob.type)).toBeInTheDocument();
    
    // Check if salary is displayed
    expect(screen.getByText(mockJob.salary)).toBeInTheDocument();
    
    // Check if description is displayed (truncated)
    expect(screen.getByText(/We are looking for an experienced Frontend Developer/)).toBeInTheDocument();
    
    // Check if requirements are displayed
    mockJob.requirements.forEach(req => {
      expect(screen.getByText(req)).toBeInTheDocument();
    });
    
    // Check if "View Details" button exists
    expect(screen.getByRole('link', { name: /View Details/i })).toHaveAttribute('href', `/jobs/${mockJob.id}`);
  });
  
  it('displays the correct category badge', () => {
    renderWithRouter(<JobCard job={mockJob} />);
    
    // Check if category badge is displayed
    expect(screen.getByText(mockJob.category)).toBeInTheDocument();
  });
});