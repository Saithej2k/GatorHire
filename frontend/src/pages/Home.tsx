import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Search, Briefcase, Bell, TrendingUp, Github, Linkedin, User, Code, Stethoscope, GraduationCap, Building2, Pencil, ChefHat, ArrowRight } from 'lucide-react';

export const Home = () => {
  const [showCategories, setShowCategories] = useState(false);
  const navigate = useNavigate();

  const handleCategoryClick = () => {
    navigate('/signup');
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col">
      {/* Hero Section */}
      <div className="bg-blue-900 text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20">
          <div className="text-center">
            <h1 className="text-5xl font-bold mb-6">
              Find Your Dream Job with <span className="text-orange-500">GatorHire</span>
            </h1>
            <p className="text-xl mb-12">
              Explore opportunities across multiple industries and take the next step in your career
            </p>
            {!showCategories ? (
              <button
                onClick={() => setShowCategories(true)}
                className="bg-orange-500 hover:bg-orange-600 text-white px-8 py-4 rounded-lg text-lg font-semibold inline-flex items-center gap-2 transition-all transform hover:scale-105"
              >
                Get Started <ArrowRight className="h-5 w-5" />
              </button>
            ) : (
              <div className="animate-fadeIn">
                <p className="text-lg mb-8">Select your preferred job category to begin</p>
                <div className="grid grid-cols-2 md:grid-cols-3 gap-4 max-w-4xl mx-auto">
                  <button
                    onClick={handleCategoryClick}
                    className="bg-white/10 hover:bg-white/20 backdrop-blur-sm p-6 rounded-lg flex flex-col items-center transition-all hover:transform hover:-translate-y-1"
                  >
                    <Code className="h-8 w-8 text-orange-500 mb-3" />
                    <span className="text-lg font-medium">Technology</span>
                  </button>
                  <button
                    onClick={handleCategoryClick}
                    className="bg-white/10 hover:bg-white/20 backdrop-blur-sm p-6 rounded-lg flex flex-col items-center transition-all hover:transform hover:-translate-y-1"
                  >
                    <Stethoscope className="h-8 w-8 text-orange-500 mb-3" />
                    <span className="text-lg font-medium">Healthcare</span>
                  </button>
                  <button
                    onClick={handleCategoryClick}
                    className="bg-white/10 hover:bg-white/20 backdrop-blur-sm p-6 rounded-lg flex flex-col items-center transition-all hover:transform hover:-translate-y-1"
                  >
                    <GraduationCap className="h-8 w-8 text-orange-500 mb-3" />
                    <span className="text-lg font-medium">Education</span>
                  </button>
                  <button
                    onClick={handleCategoryClick}
                    className="bg-white/10 hover:bg-white/20 backdrop-blur-sm p-6 rounded-lg flex flex-col items-center transition-all hover:transform hover:-translate-y-1"
                  >
                    <Building2 className="h-8 w-8 text-orange-500 mb-3" />
                    <span className="text-lg font-medium">Business</span>
                  </button>
                  <button
                    onClick={handleCategoryClick}
                    className="bg-white/10 hover:bg-white/20 backdrop-blur-sm p-6 rounded-lg flex flex-col items-center transition-all hover:transform hover:-translate-y-1"
                  >
                    <Pencil className="h-8 w-8 text-orange-500 mb-3" />
                    <span className="text-lg font-medium">Creative</span>
                  </button>
                  <button
                    onClick={handleCategoryClick}
                    className="bg-white/10 hover:bg-white/20 backdrop-blur-sm p-6 rounded-lg flex flex-col items-center transition-all hover:transform hover:-translate-y-1"
                  >
                    <ChefHat className="h-8 w-8 text-orange-500 mb-3" />
                    <span className="text-lg font-medium">Hospitality</span>
                  </button>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Features Section */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20">
        <h2 className="text-3xl font-bold text-center text-blue-900 mb-12">
          Why Choose GatorHire?
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          <div className="bg-white p-6 rounded-lg shadow-md">
            <Search className="h-12 w-12 text-orange-500 mb-4" />
            <h3 className="text-xl font-semibold text-blue-900 mb-2">
              Unified Search
            </h3>
            <p className="text-gray-600">
              Search across multiple job boards in one place with advanced filtering options.
            </p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-md">
            <Bell className="h-12 w-12 text-orange-500 mb-4" />
            <h3 className="text-xl font-semibold text-blue-900 mb-2">
              Job Alerts
            </h3>
            <p className="text-gray-600">
              Get notified about new opportunities that match your preferences.
            </p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-md">
            <TrendingUp className="h-12 w-12 text-orange-500 mb-4" />
            <h3 className="text-xl font-semibold text-blue-900 mb-2">
              Smart Recommendations
            </h3>
            <p className="text-gray-600">
              Receive personalized job suggestions based on your profile and preferences.
            </p>
          </div>
        </div>
      </div>

      {/* Team Section */}
      <div className="bg-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-center text-blue-900 mb-12">
            Meet Our Team
          </h2>
          <div className="grid md:grid-cols-3 gap-12">
            <div className="text-center">
              <div className="relative">
                <div className="w-24 h-24 bg-orange-100 rounded-full mx-auto mb-4 flex items-center justify-center">
                  <User className="h-12 w-12 text-orange-500" />
                </div>
              </div>
              <h3 className="text-xl font-semibold text-blue-900">Venkata Sai Sathvika Ande</h3>
              <p className="text-orange-500 mb-2">Front End Engineer</p>
              <div className="flex justify-center space-x-3">
                <a href="#" className="text-gray-600 hover:text-blue-900">
                  <Github className="h-5 w-5" />
                </a>
                <a href="#" className="text-gray-600 hover:text-blue-900">
                  <Linkedin className="h-5 w-5" />
                </a>
              </div>
            </div>

            <div className="text-center">
              <div className="relative">
                <div className="w-24 h-24 bg-orange-100 rounded-full mx-auto mb-4 flex items-center justify-center">
                  <User className="h-12 w-12 text-orange-500" />
                </div>
              </div>
              <h3 className="text-xl font-semibold text-blue-900">Aditya Dudugu</h3>
              <p className="text-orange-500 mb-2">Back End Engineer</p>
              <div className="flex justify-center space-x-3">
                <a href="#" className="text-gray-600 hover:text-blue-900">
                  <Github className="h-5 w-5" />
                </a>
                <a href="#" className="text-gray-600 hover:text-blue-900">
                  <Linkedin className="h-5 w-5" />
                </a>
              </div>
            </div>

            <div className="text-center">
              <div className="relative">
                <div className="w-24 h-24 bg-orange-100 rounded-full mx-auto mb-4 flex items-center justify-center">
                  <User className="h-12 w-12 text-orange-500" />
                </div>
              </div>
              <h3 className="text-xl font-semibold text-blue-900">Saithej Singhu</h3>
              <p className="text-orange-500 mb-2">Back End Engineer</p>
              <div className="flex justify-center space-x-3">
                <a href="#" className="text-gray-600 hover:text-blue-900">
                  <Github className="h-5 w-5" />
                </a>
                <a href="#" className="text-gray-600 hover:text-blue-900">
                  <Linkedin className="h-5 w-5" />
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Footer */}
      <footer className="bg-blue-900 text-white mt-auto">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
          <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
            <div className="col-span-2">
              <h3 className="text-xl font-semibold mb-4">About GatorHire</h3>
              <p className="text-gray-300">
                GatorHire is a comprehensive job-hunting platform that combines listings from top job boards,
                making your job search more efficient and effective.
              </p>
            </div>
            <div>
              <h3 className="text-xl font-semibold mb-4">Quick Links</h3>
              <ul className="space-y-2">
                <li><a href="#" className="text-gray-300 hover:text-orange-500">Home</a></li>
                <li><a href="#" className="text-gray-300 hover:text-orange-500">About Us</a></li>
                <li><a href="#" className="text-gray-300 hover:text-orange-500">Contact</a></li>
                <li><a href="#" className="text-gray-300 hover:text-orange-500">Privacy Policy</a></li>
              </ul>
            </div>
            <div>
              <h3 className="text-xl font-semibold mb-4">Contact Us</h3>
              <ul className="space-y-2">
                <li className="text-gray-300">University of Florida</li>
                <li className="text-gray-300">Gainesville, FL 32611</li>
                <li className="text-gray-300">support@gatorhire.com</li>
              </ul>
            </div>
          </div>
          <div className="mt-8 pt-8 border-t border-blue-800">
            <div className="text-center text-gray-300">
              <p>&copy; {new Date().getFullYear()} GatorHire. All rights reserved.</p>
              <p className="mt-2 text-sm">
                A project by University of Florida students. Powered by React and Go.
              </p>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
};