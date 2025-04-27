import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Briefcase, Users, Building, GraduationCap, Palette, Coffee, CheckCircle, ArrowRight, ArrowLeft } from 'lucide-react';
import { JobCategory } from '../frontend/types/job';
import CategorySelector from '../frontend/components/home/CategorySelector';

// Quiz questions for career matching
const quizQuestions = [
  {
    question: "What type of work environment do you prefer?",
    options: [
      { text: "Fast-paced startup", categories: ["Technology", "Creative"] },
      { text: "Established corporation", categories: ["Business", "Technology"] },
      { text: "Community-focused organization", categories: ["Healthcare", "Education"] },
      { text: "Flexible/remote work", categories: ["Technology", "Creative"] }
    ]
  },
  {
    question: "Which skills do you enjoy using the most?",
    options: [
      { text: "Technical and analytical", categories: ["Technology", "Business"] },
      { text: "Creative and design", categories: ["Creative", "Hospitality"] },
      { text: "Communication and people skills", categories: ["Healthcare", "Education", "Hospitality"] },
      { text: "Organization and management", categories: ["Business", "Hospitality"] }
    ]
  },
  {
    question: "What's most important to you in a job?",
    options: [
      { text: "High salary and benefits", categories: ["Technology", "Business", "Healthcare"] },
      { text: "Work-life balance", categories: ["Education", "Creative"] },
      { text: "Making a difference", categories: ["Healthcare", "Education"] },
      { text: "Learning and growth", categories: ["Technology", "Business", "Creative"] }
    ]
  }
];

const HomePage: React.FC = () => {
  const [selectedCategory, setSelectedCategory] = useState<JobCategory | ''>('');
  const navigate = useNavigate();
  
  // Quiz state
  const [showQuiz, setShowQuiz] = useState(false);
  const [currentQuestion, setCurrentQuestion] = useState(0);
  const [answers, setAnswers] = useState<number[]>([]);
  const [quizResult, setQuizResult] = useState<JobCategory | null>(null);

  const handleCategorySelect = (category: JobCategory) => {
    setSelectedCategory(category);
  };

  const startQuiz = () => {
    setShowQuiz(true);
    setCurrentQuestion(0);
    setAnswers([]);
    setQuizResult(null);
  };

  const handleAnswer = (answerIndex: number) => {
    const newAnswers = [...answers, answerIndex];
    setAnswers(newAnswers);
    
    if (currentQuestion < quizQuestions.length - 1) {
      setCurrentQuestion(currentQuestion + 1);
    } else {
      // Calculate result
      const categoryCounts: Record<string, number> = {};
      
      newAnswers.forEach((answerIdx, questionIdx) => {
        const selectedCategories = quizQuestions[questionIdx].options[answerIdx].categories;
        selectedCategories.forEach(category => {
          categoryCounts[category] = (categoryCounts[category] || 0) + 1;
        });
      });
      
      // Find category with highest count
      let maxCount = 0;
      let resultCategory: JobCategory = 'Technology';
      
      Object.entries(categoryCounts).forEach(([category, count]) => {
        if (count > maxCount) {
          maxCount = count;
          resultCategory = category as JobCategory;
        }
      });
      
      setQuizResult(resultCategory);
    }
  };

  const resetQuiz = () => {
    setShowQuiz(false);
    setCurrentQuestion(0);
    setAnswers([]);
    setQuizResult(null);
  };

  const navigateToJobsByResult = () => {
    if (quizResult) {
      navigate(`/jobs?category=${quizResult}`);
    }
  };

  const goToPreviousQuestion = () => {
    if (currentQuestion > 0) {
      setCurrentQuestion(currentQuestion - 1);
      setAnswers(answers.slice(0, -1));
    }
  };

  return (
    <div className="flex flex-col min-h-screen">
      {/* Hero Section */}
      <section className="bg-gradient-to-r from-orange-500 to-orange-600 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <h1 className="text-4xl md:text-5xl font-bold mb-6">Find Your Dream Job with GatorHire</h1>
            <p className="text-xl mb-8">Connect with top employers and discover opportunities that match your skills and aspirations.</p>
            
            <div className="max-w-3xl mx-auto">
              {!showQuiz && !quizResult && (
                <button
                  onClick={startQuiz}
                  className="bg-blue-600 hover:bg-blue-700 text-white px-8 py-4 rounded-md font-medium transition duration-300 text-lg shadow-lg transform hover:scale-105"
                >
                  Take Our Career Match Quiz
                </button>
              )}
              
              {showQuiz && !quizResult && (
                <div className="bg-white text-gray-800 p-6 rounded-lg shadow-lg">
                  <div className="mb-4 flex justify-between items-center">
                    <span className="text-sm font-medium text-gray-500">Question {currentQuestion + 1} of {quizQuestions.length}</span>
                    {currentQuestion > 0 && (
                      <button 
                        onClick={goToPreviousQuestion}
                        className="text-blue-600 hover:text-blue-800 flex items-center text-sm"
                      >
                        <ArrowLeft className="h-4 w-4 mr-1" />
                        Previous
                      </button>
                    )}
                  </div>
                  
                  <h3 className="text-xl font-bold mb-4">{quizQuestions[currentQuestion].question}</h3>
                  
                  <div className="space-y-3">
                    {quizQuestions[currentQuestion].options.map((option, idx) => (
                      <button
                        key={idx}
                        onClick={() => handleAnswer(idx)}
                        className="w-full text-left p-4 border border-gray-200 rounded-md hover:bg-blue-50 hover:border-blue-300 transition duration-200"
                      >
                        {option.text}
                      </button>
                    ))}
                  </div>
                </div>
              )}
              
              {quizResult && (
                <div className="bg-white text-gray-800 p-6 rounded-lg shadow-lg">
                  <div className="flex justify-center mb-4">
                    <CheckCircle className="h-16 w-16 text-green-500" />
                  </div>
                  <h3 className="text-2xl font-bold mb-2">Your Career Match: {quizResult}</h3>
                  <p className="mb-6 text-gray-600">Based on your preferences, we think you'd excel in {quizResult} roles!</p>
                  
                  <div className="flex flex-col sm:flex-row gap-4 justify-center">
                    <button
                      onClick={navigateToJobsByResult}
                      className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-md font-medium transition duration-300 flex items-center justify-center"
                    >
                      Browse {quizResult} Jobs
                      <ArrowRight className="h-5 w-5 ml-2" />
                    </button>
                    <button
                      onClick={resetQuiz}
                      className="bg-gray-200 hover:bg-gray-300 text-gray-800 px-6 py-3 rounded-md font-medium transition duration-300"
                    >
                      Retake Quiz
                    </button>
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>
      </section>

      {/* Category Selection Section */}
      <section className="py-16 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">Explore Job Categories</h2>
            <p className="text-gray-600 max-w-2xl mx-auto">
              Discover opportunities in various industries and find the perfect role that matches your skills and interests.
            </p>
          </div>
          
          <CategorySelector onCategorySelect={handleCategorySelect} />
        </div>
      </section>

      {/* How It Works */}
      <section className="py-16 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-center mb-12">How GatorHire Works</h2>
          
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="text-center">
              <div className="bg-orange-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <span className="text-orange-600 text-2xl font-bold">1</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">Create Your Profile</h3>
              <p className="text-gray-600">Sign up and build your professional profile to showcase your skills and experience.</p>
            </div>
            
            <div className="text-center">
              <div className="bg-orange-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <span className="text-orange-600 text-2xl font-bold">2</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">Discover Opportunities</h3>
              <p className="text-gray-600">Browse job listings or receive personalized recommendations based on your profile.</p>
            </div>
            
            <div className="text-center">
              <div className="bg-orange-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <span className="text-orange-600 text-2xl font-bold">3</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">Apply with Ease</h3>
              <p className="text-gray-600">Submit applications with just a few clicks and track your application status.</p>
            </div>
          </div>
        </div>
      </section>

      {/* Featured Jobs Section */}
      <section className="py-16 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center mb-8">
            <h2 className="text-3xl font-bold">Featured Jobs</h2>
            <Link to="/jobs" className="text-blue-600 hover:text-blue-800 font-medium">
              View All Jobs →
            </Link>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {/* Featured Job 1 */}
            <div className="bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow p-6">
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-lg font-semibold text-gray-900">Senior Frontend Developer</h3>
                <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded-full">Technology</span>
              </div>
              <div className="mb-4">
                <div className="flex items-center text-gray-600 mb-1">
                  <Briefcase className="h-4 w-4 mr-2" />
                  <span>TechCorp</span>
                </div>
                <div className="flex items-center text-gray-600">
                  <Building className="h-4 w-4 mr-2" />
                  <span>San Francisco, CA</span>
                </div>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-900 font-medium">$120K - $150K</span>
                <Link to="/jobs/1" className="text-blue-600 hover:text-blue-800">
                  View Details →
                </Link>
              </div>
            </div>
            
            {/* Featured Job 2 */}
            <div className="bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow p-6">
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-lg font-semibold text-gray-900">Registered Nurse</h3>
                <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full">Healthcare</span>
              </div>
              <div className="mb-4">
                <div className="flex items-center text-gray-600 mb-1">
                  <Briefcase className="h-4 w-4 mr-2" />
                  <span>City Hospital</span>
                </div>
                <div className="flex items-center text-gray-600">
                  <Building className="h-4 w-4 mr-2" />
                  <span>Chicago, IL</span>
                </div>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-900 font-medium">$75K - $95K</span>
                <Link to="/jobs/6" className="text-blue-600 hover:text-blue-800">
                  View Details →
                </Link>
              </div>
            </div>
            
            {/* Featured Job 3 */}
            <div className="bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow p-6">
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-lg font-semibold text-gray-900">Marketing Manager</h3>
                <span className="bg-purple-100 text-purple-800 text-xs font-medium px-2.5 py-0.5 rounded-full">Business</span>
              </div>
              <div className="mb-4">
                <div className="flex items-center text-gray-600 mb-1">
                  <Briefcase className="h-4 w-4 mr-2" />
                  <span>Brand Solutions</span>
                </div>
                <div className="flex items-center text-gray-600">
                  <Building className="h-4 w-4 mr-2" />
                  <span>Los Angeles, CA</span>
                </div>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-900 font-medium">$85K - $110K</span>
                <Link to="/jobs/11" className="text-blue-600 hover:text-blue-800">
                  View Details →
                </Link>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Testimonials Section */}
      <section className="py-16 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-center mb-12">What Our Users Say</h2>
          
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {/* Testimonial 1 */}
            <div className="bg-white p-6 rounded-lg shadow-md">
              <div className="flex items-center mb-4">
                <div className="h-12 w-12 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 font-bold text-xl">
                  JD
                </div>
                <div className="ml-4">
                  <h3 className="font-semibold">John Doe</h3>
                  <p className="text-gray-600 text-sm">Software Engineer</p>
                </div>
              </div>
              <p className="text-gray-700">
                "GatorHire made my job search incredibly easy. I found my dream position at a tech company within just two weeks of signing up!"
              </p>
            </div>
            
            {/* Testimonial 2 */}
            <div className="bg-white p-6 rounded-lg shadow-md">
              <div className="flex items-center mb-4">
                <div className="h-12 w-12 rounded-full bg-green-100 flex items-center justify-center text-green-600 font-bold text-xl">
                  SJ
                </div>
                <div className="ml-4">
                  <h3 className="font-semibold">Sarah Johnson</h3>
                  <p className="text-gray-600 text-sm">Marketing Specialist</p>
                </div>
              </div>
              <p className="text-gray-700">
                "The personalized job recommendations were spot on. I received notifications for positions that perfectly matched my skills and experience."
              </p>
            </div>
            
            {/* Testimonial 3 */}
            <div className="bg-white p-6 rounded-lg shadow-md">
              <div className="flex items-center mb-4">
                <div className="h-12 w-12 rounded-full bg-purple-100 flex items-center justify-center text-purple-600 font-bold text-xl">
                  MT
                </div>
                <div className="ml-4">
                  <h3 className="font-semibold">Michael Thompson</h3>
                  <p className="text-gray-600 text-sm">Healthcare Professional</p>
                </div>
              </div>
              <p className="text-gray-700">
                "As someone in the healthcare industry, I appreciated how easy it was to filter jobs by specialty. Found a great position at a top hospital!"
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="bg-blue-600 text-white py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl font-bold mb-6">Ready to Start Your Job Search?</h2>
          <p className="text-xl mb-8 max-w-3xl mx-auto">Join thousands of professionals who have found their dream jobs through GatorHire.</p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link to="/jobs" className="bg-white text-blue-600 hover:bg-gray-100 px-6 py-3 rounded-md font-medium transition duration-300">
              Browse Jobs
            </Link>
            <Link to="/signup" className="bg-orange-600 hover:bg-orange-700 text-white px-6 py-3 rounded-md font-medium transition duration-300">
              Create Profile
            </Link>
          </div>
        </div>
      </section>
    </div>
  );
};

export default HomePage;