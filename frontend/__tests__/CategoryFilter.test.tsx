import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import CategoryFilter from '../components/jobs/CategoryFilter';
import { JobCategory } from '../types/job';

describe('CategoryFilter Component', () => {
  const mockCategoryCounts: Record<JobCategory, number> = {
    'All': 15,
    'Technology': 5,
    'Healthcare': 2,
    'Education': 2,
    'Business': 2,
    'Creative': 2,
    'Hospitality': 2
  };
  
  const mockOnCategoryChange = vi.fn();
  
  beforeEach(() => {
    mockOnCategoryChange.mockClear();
  });
  
  it('renders all category buttons', () => {
    render(
      <CategoryFilter 
        selectedCategory="All" 
        onCategoryChange={mockOnCategoryChange} 
        categoryCounts={mockCategoryCounts} 
      />
    );
    
    // Check if all category buttons are rendered
    expect(screen.getByText('All')).toBeInTheDocument();
    expect(screen.getByText('Technology')).toBeInTheDocument();
    expect(screen.getByText('Healthcare')).toBeInTheDocument();
    expect(screen.getByText('Education')).toBeInTheDocument();
    expect(screen.getByText('Business')).toBeInTheDocument();
    expect(screen.getByText('Creative')).toBeInTheDocument();
    expect(screen.getByText('Hospitality')).toBeInTheDocument();
  });
  
  it('displays the correct count for each category', () => {
    render(
      <CategoryFilter 
        selectedCategory="All" 
        onCategoryChange={mockOnCategoryChange} 
        categoryCounts={mockCategoryCounts} 
      />
    );
    
    // Check if category counts are displayed correctly
    expect(screen.getByText('15')).toBeInTheDocument(); // All
    expect(screen.getByText('5')).toBeInTheDocument(); // Technology
    expect(screen.getAllByText('2').length).toBe(5); // Other categories
  });
  
  it('highlights the selected category', () => {
    render(
      <CategoryFilter 
        selectedCategory="Technology" 
        onCategoryChange={mockOnCategoryChange} 
        categoryCounts={mockCategoryCounts} 
      />
    );
    
    // The Technology button should have a different style when selected
    const technologyButton = screen.getByText('Technology').closest('button');
    expect(technologyButton).toHaveClass('text-blue-800');
  });
  
  it('calls onCategoryChange when a category is clicked', () => {
    render(
      <CategoryFilter 
        selectedCategory="All" 
        onCategoryChange={mockOnCategoryChange} 
        categoryCounts={mockCategoryCounts} 
      />
    );
    
    // Click on the Technology category
    fireEvent.click(screen.getByText('Technology'));
    
    // Check if onCategoryChange was called with the correct category
    expect(mockOnCategoryChange).toHaveBeenCalledWith('Technology');
  });
});