export interface User {
    id: string;
    email: string;
    name: string;
  }
  
  export interface AuthState {
    user: User | null;
    isAuthenticated: boolean;
  }
  
  export interface Job {
    id: string;
    title: string;
    company: string;
    location: string;
    salary: string;
    type: string;
    description: string;
    requirements: string[];
    postedDate: Date;
  }
  
  export interface JobCategory {
    id: string;
    name: string;
    description: string;
  }
  
  export interface JobApplication {
    id: string;
    jobId: string;
    userId: string;
    status: 'pending' | 'reviewed' | 'accepted' | 'rejected';
    appliedDate: Date;
    resumeUrl: string;
    coverLetter?: string;
  }
  
  export interface ApiError extends Error {
    code: string;
    status: number;
  }
  
  export type ApiResponse<T> = {
    data?: T;
    error?: ApiError;
  };