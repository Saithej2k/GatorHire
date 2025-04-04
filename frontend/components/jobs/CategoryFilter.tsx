import React from 'react';
import { Briefcase, Users, GraduationCap, Building, Palette, Coffee } from 'lucide-react';
import { JobCategory } from '../../types/job';

interface CategoryFilterProps {
  selectedCategory: JobCategory;
  onCategoryChange: (category: JobCategory) => void;
  categoryCounts: Record<JobCategory, number>;
}

const CategoryFilter: React.FC<CategoryFilterProps> = ({
  selectedCategory,
  onCategoryChange,
  categoryCounts,
}) => {
  const categories: { name: JobCategory; icon: React.ReactNode; color: string; bgColor: string }[] = [
    { name: 'All', icon: <Briefcase />, color: 'text-gray-800', bgColor: 'bg-gray-100' },
    { name: 'Technology', icon: <Briefcase />, color: 'text-blue-800', bgColor: 'bg-blue-100' },
    { name: 'Healthcare', icon: <Users />, color: 'text-green-800', bgColor: 'bg-green-100' },
    { name: 'Education', icon: <GraduationCap />, color: 'text-orange-800', bgColor: 'bg-orange-100' },
    { name: 'Business', icon: <Building />, color: 'text-purple-800', bgColor: 'bg-purple-100' },
    { name: 'Creative', icon: <Palette />, color: 'text-pink-800', bgColor: 'bg-pink-100' },
    { name: 'Hospitality', icon: <Coffee />, color: 'text-yellow-800', bgColor: 'bg-yellow-100' },
  ];

  return (
    <div className="mb-6 overflow-x-auto pb-2">
      <div className="flex space-x-2 min-w-max">
        {categories.map((category) => {
          const isSelected = selectedCategory === category.name;
          return (
            <button
              key={category.name}
              onClick={() => onCategoryChange(category.name)}
              className={`flex items-center px-4 py-2 rounded-full text-sm font-medium transition-colors ${
                isSelected
                  ? `${category.bgColor} ${category.color} border-2 border-current`
                  : 'bg-white text-gray-600 hover:bg-gray-100'
              }`}
            >
              <span className="mr-2">{category.icon}</span>
              {category.name}
              {categoryCounts[category.name] !== undefined && (
                <span className={`ml-2 ${isSelected ? 'bg-white' : category.bgColor} text-gray-600 rounded-full px-2 py-0.5 text-xs`}>
                  {categoryCounts[category.name] || 0}
                </span>
              )}
            </button>
          );
        })}
      </div>
    </div>
  );
};

export default CategoryFilter;