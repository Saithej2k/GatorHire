import React from 'react';
import { Link } from 'react-router-dom';
import { Briefcase } from 'lucide-react';

export const Navbar = () => {
  return (
    <nav className="bg-white shadow-md">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link to="/" className="flex items-center">
              <Briefcase className="h-8 w-8 text-orange-500" />
              <span className="ml-2 text-2xl font-bold text-blue-900">GatorHire</span>
            </Link>
          </div>
          <div className="flex items-center space-x-4">
            <Link to="/login" className="text-blue-900 hover:text-orange-500 px-3 py-2">
              Login
            </Link>
            <Link
              to="/signup"
              className="bg-orange-500 text-white px-4 py-2 rounded-md hover:bg-orange-600"
            >
              Sign Up
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
};