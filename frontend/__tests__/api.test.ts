import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import axios from 'axios';
import { fetchJobs, fetchJobById, submitApplication } from '../services/api';
import { ApplicationFormData } from '../types/application';

// Mock axios
vi.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;

describe('API Service', () => {
  beforeEach(() => {
    // Clear all mocks before each test
    vi.clearAllMocks();
    
    // Setup axios create mock
    mockedAxios.create = vi.fn(() => mockedAxios);
  });

  afterEach(() => {
    vi.resetAllMocks();
  });

  describe('fetchJobs', () => {
    it('should fetch all jobs successfully', async () => {
      // Mock data
      const mockJobs = [
        { id: '1', title: 'Job 1', category: 'Technology' },
        { id: '2', title: 'Job 2', category: 'Healthcare' }
      ];
      
      // Setup mock response
      mockedAxios.get.mockResolvedValueOnce({ data: mockJobs });
      
      // Call the function
      const result = await fetchJobs();
      
      // Assertions
      expect(mockedAxios.get).toHaveBeenCalledWith('/jobs');
      expect(result).toEqual(mockJobs);
    });

    it('should handle errors when fetching jobs fails', async () => {
      // Setup mock error
      const errorMessage = 'Network Error';
      mockedAxios.get.mockRejectedValueOnce(new Error(errorMessage));
      
      // Call the function and expect it to throw
      await expect(fetchJobs()).rejects.toThrow();
      
      // Assertions
      expect(mockedAxios.get).toHaveBeenCalledWith('/jobs');
    });
  });

  describe('fetchJobById', () => {
    it('should fetch a job by ID successfully', async () => {
      // Mock data
      const mockJob = { id: '1', title: 'Job 1', category: 'Technology' };
      const jobId = '1';
      
      // Setup mock response
      mockedAxios.get.mockResolvedValueOnce({ data: mockJob });
      
      // Call the function
      const result = await fetchJobById(jobId);
      
      // Assertions
      expect(mockedAxios.get).toHaveBeenCalledWith(`/jobs/${jobId}`);
      expect(result).toEqual(mockJob);
    });

    it('should throw an error if job ID is not provided', async () => {
      // Call the function with undefined ID and expect it to throw
      await expect(fetchJobById(undefined)).rejects.toThrow('Job ID is required');
      
      // Assertions
      expect(mockedAxios.get).not.toHaveBeenCalled();
    });
  });

  describe('submitApplication', () => {
    it('should submit an application successfully', async () => {
      // Mock data
      const jobId = '1';
      const formData: ApplicationFormData = {
        fullName: 'John Doe',
        email: 'john@example.com',
        phone: '123-456-7890',
        coverLetter: 'I am interested in this position',
        resumeFile: null,
        linkedIn: '',
        portfolio: '',
        heardFrom: ''
      };
      const mockResponse = { success: true, applicationId: '123' };
      
      // Setup mock response
      mockedAxios.post.mockResolvedValueOnce({ data: mockResponse });
      
      // Call the function
      const result = await submitApplication(jobId, formData);
      
      // Assertions
      expect(mockedAxios.post).toHaveBeenCalledWith('/applications', {
        jobId,
        ...formData
      });
      expect(result).toEqual(mockResponse);
    });

    it('should throw an error if job ID is not provided', async () => {
      // Mock data
      const formData: ApplicationFormData = {
        fullName: 'John Doe',
        email: 'john@example.com',
        phone: '123-456-7890',
        coverLetter: 'I am interested in this position',
        resumeFile: null,
        linkedIn: '',
        portfolio: '',
        heardFrom: ''
      };
      
      // Call the function with undefined ID and expect it to throw
      await expect(submitApplication(undefined, formData)).rejects.toThrow('Job ID is required');
      
      // Assertions
      expect(mockedAxios.post).not.toHaveBeenCalled();
    });
  });
});