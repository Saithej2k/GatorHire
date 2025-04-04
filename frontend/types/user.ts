/**
 * Types related to users
 */

export interface User {
  id: string;
  email: string;
  fullName: string;
  title?: string;
  location?: string;
  bio?: string;
  skills?: string[];
  role: 'user' | 'admin';
}

export interface UserProfile {
  fullName: string;
  email: string;
  phone: string;
  location: string;
  title: string;
  bio: string;
  skills: string[];
  experience: Experience[];
  education: Education[];
}

export interface Experience {
  id: number;
  company: string;
  title: string;
  startDate: string;
  endDate: string;
  description: string;
}

export interface Education {
  id: number;
  institution: string;
  degree: string;
  year: string;
}