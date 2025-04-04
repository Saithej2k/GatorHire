import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import SearchInput from '../components/ui/SearchInput';

describe('SearchInput Component', () => {
  it('renders with placeholder text', () => {
    const mockOnChange = vi.fn();
    render(<SearchInput value="" onChange={mockOnChange} placeholder="Test placeholder" />);
    
    // Check if input with placeholder exists
    const input = screen.getByPlaceholderText('Test placeholder');
    expect(input).toBeInTheDocument();
  });
  
  it('displays the current value', () => {
    const mockOnChange = vi.fn();
    render(<SearchInput value="test value" onChange={mockOnChange} />);
    
    // Check if input has the correct value
    const input = screen.getByRole('textbox') as HTMLInputElement;
    expect(input.value).toBe('test value');
  });
  
  it('calls onChange when input value changes', () => {
    const mockOnChange = vi.fn();
    render(<SearchInput value="" onChange={mockOnChange} />);
    
    // Get the input element
    const input = screen.getByRole('textbox');
    
    // Simulate typing in the input
    fireEvent.change(input, { target: { value: 'new value' } });
    
    // Check if onChange was called with the correct event
    expect(mockOnChange).toHaveBeenCalledTimes(1);
    expect(mockOnChange.mock.calls[0][0].target.value).toBe('new value');
  });
  
  it('applies custom className', () => {
    const mockOnChange = vi.fn();
    render(<SearchInput value="" onChange={mockOnChange} className="custom-class" />);
    
    // Check if the custom class is applied to the container
    const container = screen.getByRole('textbox').parentElement;
    expect(container).toHaveClass('custom-class');
  });
});