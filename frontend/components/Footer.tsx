import React from 'react';
import { Link } from 'react-router-dom';
import { Briefcase, Github, Linkedin, Mail } from 'lucide-react';

const Footer: React.FC = () => {
  return (
    <footer className="bg-blue-800 text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div>
            <div className="flex items-center mb-4">
              <Briefcase className="h-6 w-6 mr-2" />
              <span className="text-xl font-bold">GatorHire</span>
            </div>
            <p className="text-gray-300">
              Connecting talented professionals with great opportunities.
            </p>
          </div>
          
          <div>
            <h3 className="text-lg font-semibold mb-4">Quick Links</h3>
            <ul className="space-y-2">
              <li>
                <Link to="/" className="text-gray-300 hover:text-white">Home</Link>
              </li>
              <li>
                <Link to="/jobs" className="text-gray-300 hover:text-white">Browse Jobs</Link>
              </li>
              <li>
                <Link to="/profile" className="text-gray-300 hover:text-white">Profile</Link>
              </li>
            </ul>
          </div>
          
          <div>
            <h3 className="text-lg font-semibold mb-4">Connect With Us</h3>
            <div className="flex space-x-4">
              <a href="https://github.com/Saithej2k/GatorHire" className="text-gray-300 hover:text-white">
                <Github className="h-6 w-6" />
              </a>
              <a href="#" className="text-gray-300 hover:text-white">
                <Linkedin className="h-6 w-6" />
              </a>
              <a href="mailto:contact@gatorhire.com" className="text-gray-300 hover:text-white">
                <Mail className="h-6 w-6" />
              </a>
            </div>
          </div>
        </div>
        
        <div className="mt-8 pt-8 border-t border-gray-700 text-center text-gray-300">
          <p>&copy; {new Date().getFullYear()} GatorHire. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;