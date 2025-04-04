import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Briefcase, Users, GraduationCap, Building, Palette, Coffee, ChevronRight } from 'lucide-react';
import { JobCategory } from '../../types/job';

interface CategorySelectorProps {
  onCategorySelect: (category: JobCategory) => void;
}

const CategorySelector: React.FC<CategorySelectorProps> = ({ onCategorySelect }) => {
  const navigate = useNavigate();
  const [hoveredCategory, setHoveredCategory] = useState<JobCategory | null>(null);
  
  const categories = [
    {
      name: 'Technology' as JobCategory,
      icon: <Briefcase className="h-8 w-8 text-blue-600" />,
      color: 'from-blue-500 to-blue-600',
      hoverColor: 'from-blue-600 to-blue-700',
      description: 'Software development, IT, data science, and more',
      jobs: ['Frontend Developer', 'Data Scientist', 'DevOps Engineer']
    },
    {
      name: 'Healthcare' as JobCategory,
      icon: <Users className="h-8 w-8 text-green-600" />,
      color: 'from-green-500 to-green-600',
      hoverColor: 'from-green-600 to-green-700',
      description: 'Medical, nursing, therapy, and healthcare administration',
      jobs: ['Registered Nurse', 'Physical Therapist', 'Medical Technician']
    },
    {
      name: 'Education' as JobCategory,
      icon: <GraduationCap className="h-8 w-8 text-orange-600" />,
      color: 'from-orange-500 to-orange-600',
      hoverColor: 'from-orange-600 to-orange-700',
      description: 'Teaching, administration, and educational technology',
      jobs: ['Teacher', 'School Principal', 'Education Consultant']
    },
    {
      name: 'Business' as JobCategory,
      icon: <Building className="h-8 w-8 text-purple-600" />,
      color: 'from-purple-500 to-purple-600',
      hoverColor: 'from-purple-600 to-purple-700',
      description: 'Finance, marketing, sales, and management',
      jobs: ['Financial Analyst', 'Marketing Manager', 'Business Consultant']
    },
    {
      name: 'Creative' as JobCategory,
      icon: <Palette className="h-8 w-8 text-pink-600" />,
      color: 'from-pink-500 to-pink-600',
      hoverColor: 'from-pink-600 to-pink-700',
      description: 'Design, writing, media, and the arts',
      jobs: ['Graphic Designer', 'Content Writer', 'UX/UI Designer']
    },
    {
      name: 'Hospitality' as JobCategory,
      icon: <Coffee className="h-8 w-8 text-yellow-600" />,
      color: 'from-yellow-500 to-yellow-600',
      hoverColor: 'from-yellow-600 to-yellow-700',
      description: 'Hotels, restaurants, tourism, and event management',
      jobs: ['Hotel Manager', 'Executive Chef', 'Event Coordinator']
    }
  ];

  const handleCategoryClick = (category: JobCategory) => {
    onCategorySelect(category);
    // Navigate to jobs page with category filter
    navigate(`/jobs?category=${encodeURIComponent(category)}`);
  };

  return (
    <div className="w-full max-w-5xl mx-auto">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {categories.map((category) => (
          <div
            key={category.name}
            className="relative overflow-hidden rounded-xl shadow-lg transition-all duration-300 transform hover:-translate-y-1 hover:shadow-xl cursor-pointer"
            onMouseEnter={() => setHoveredCategory(category.name)}
            onMouseLeave={() => setHoveredCategory(null)}
            onClick={() => handleCategoryClick(category.name)}
          >
            <div className={`absolute inset-0 bg-gradient-to-br ${category.color} opacity-90 transition-all duration-300 ${hoveredCategory === category.name ? 'scale-110' : 'scale-100'}`}></div>
            
            <div className="relative p-6 h-full flex flex-col justify-between z-10">
              <div>
                <div className="bg-white/20 backdrop-blur-sm p-3 rounded-full w-fit mb-4">
                  {category.icon}
                </div>
                <h3 className="text-xl font-bold text-white mb-2">{category.name}</h3>
                <p className="text-white/90 mb-4">{category.description}</p>
              </div>
              
              <div>
                <div className="space-y-1 mb-4">
                  {category.jobs.map((job, index) => (
                    <div key={index} className="flex items-center text-white/80">
                      <div className="w-1 h-1 bg-white/80 rounded-full mr-2"></div>
                      <span>{job}</span>
                    </div>
                  ))}
                </div>
                
                <button className="flex items-center text-white font-medium group">
                  <span>Explore {category.name} Jobs</span>
                  <ChevronRight className="h-4 w-4 ml-1 transition-transform group-hover:translate-x-1" />
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default CategorySelector;