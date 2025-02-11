import { User, Job, JobApplication, ApiError, ApiResponse } from '../types';

// const API_URL = '/api';
const API_URL = 'http://localhost:5173/api';  // Use the correct backend URL and port


async function handleResponse<T>(response: Response): Promise<ApiResponse<T>> {
  if (!response.ok) {
    const error: ApiError = {
      name: 'ApiError',
      message: 'An error occurred',
      code: 'UNKNOWN_ERROR',
      status: response.status
    };

    try {
      const data = await response.json();
      error.message = data.message || error.message;
      error.code = data.code || error.code;
    } catch {
      // Use default error if response is not JSON
    }

    return { error };
  }

  const data = await response.json();
  return { data };
}

export const api = {
  auth: {
    login: async (email: string, password: string): Promise<ApiResponse<User>> => {
      const response = await fetch(`${API_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });
      return handleResponse<User>(response);
    },
    signup: async (email: string, password: string, name: string): Promise<ApiResponse<User>> => {
      const response = await fetch(`${API_URL}/auth/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password, name }),
      });
      return handleResponse<User>(response);
    },
    logout: async (): Promise<ApiResponse<void>> => {
      const response = await fetch(`${API_URL}/auth/logout`, {
        method: 'POST',
      });
      return handleResponse<void>(response);
    },
  },
  jobs: {
    search: async (query: string): Promise<ApiResponse<Job[]>> => {
      const response = await fetch(`${API_URL}/jobs/search?q=${encodeURIComponent(query)}`);
      return handleResponse<Job[]>(response);
    },
    getAll: async (): Promise<ApiResponse<Job[]>> => {
      const response = await fetch(`${API_URL}/jobs`);
      return handleResponse<Job[]>(response);
    },
    getById: async (id: string): Promise<ApiResponse<Job>> => {
      const response = await fetch(`${API_URL}/jobs/${id}`);
      return handleResponse<Job>(response);
    },
    apply: async (jobId: string, application: Omit<JobApplication, 'id' | 'jobId' | 'userId' | 'status' | 'appliedDate'>): Promise<ApiResponse<JobApplication>> => {
      const response = await fetch(`${API_URL}/jobs/${jobId}/apply`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(application),
      });
      return handleResponse<JobApplication>(response);
    },
  },
};