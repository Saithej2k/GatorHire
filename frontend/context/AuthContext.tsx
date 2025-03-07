import React, { createContext, useState, useContext, useEffect } from 'react';
import { User } from '../types/user';
import { login as apiLogin, adminLogin, register as apiRegister } from '../services/api';

interface AuthContextType {
  isAuthenticated: boolean;
  user: User | null;
  login: (email: string, password: string, isAdmin?: boolean) => Promise<boolean>;
  signup: (email: string, password: string, fullName: string) => Promise<boolean>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
  user: null,
  login: async () => false,
  signup: async () => false,
  logout: () => {},
});

export const useAuth = () => useContext(AuthContext);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [user, setUser] = useState<User | null>(null);

  // Check if user is already logged in
  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    const token = localStorage.getItem('token');
    
    if (storedUser && token) {
      try {
        setUser(JSON.parse(storedUser));
        setIsAuthenticated(true);
      } catch (error) {
        localStorage.removeItem('user');
        localStorage.removeItem('token');
      }
    }
  }, []);

  const login = async (email: string, password: string, isAdmin: boolean = false): Promise<boolean> => {
    try {
      console.log('Attempting login with:', email);
      let response;
  
      if (isAdmin) {
        response = await adminLogin(email, password);
      } else {
        response = await apiLogin(email, password);
      }

      console.log('response', response);
  
      // ✅ Ensure response structure is valid
      if (!response || !response.success || !response.user || !response.token) {
        console.error('Login failed:', response?.error || 'Invalid credentials');
        return false; // Return false if login fails
      }
  
      // Store user data and token in localStorage
      localStorage.setItem('token', response.token);
      localStorage.setItem('user', JSON.stringify(response.user));
  
      setUser(response.user);
      setIsAuthenticated(true);
  
      console.log('Login successful!');
      return true;
    } catch (error: any) {
      console.error('Login error:', error);
  
      // ✅ If using `fetch`, check error response
      if (error.response && error.response.status === 401) {
        console.error('Unauthorized: Invalid credentials');
      }
      
      return false;
    }
  };
  

  // const login = async (email: string, password: string, isAdmin: boolean = false): Promise<boolean> => {
  //   try {
  //     console.log('Attempting login with:', email);
  //     let response;
      
  //     if (isAdmin) {
  //       response = await adminLogin(email, password);
  //     } else {
  //       response = await apiLogin(email, password);
  //     }
      
  //     if (response.success && response.user && response.token) {
  //       // Store user data and token in localStorage
  //       localStorage.setItem('token', response.token);
  //       localStorage.setItem('user', JSON.stringify(response.user));
        
  //       setUser(response.user);
  //       setIsAuthenticated(true);
        
  //       console.log('Login successful!');
  //       return true;
  //     } else {
  //       console.error('Login failed:', response.error);
  //       return false;
  //     }
  //   } catch (error) {
  //     console.error('Login error:', error);
  //     return false;
  //   }
  // };

  const signup = async (email: string, password: string, fullName: string): Promise<boolean> => {
    try {
      console.log('Submitting registration form:', { email, fullName });
      const response = await apiRegister(email, password, fullName);
      
      if (response.success && response.user && response.token) {
        // Store user data and token in localStorage
        localStorage.setItem('token', response.token);
        localStorage.setItem('user', JSON.stringify(response.user));
        
        setUser(response.user);
        setIsAuthenticated(true);
        
        console.log('Registration successful!');
        return true;
      } else {
        console.error('Registration failed:', response.error);
        return false;
      }
    } catch (error) {
      console.error('Registration error:', error);
      return false;
    }
  };

  const logout = () => {
    // Remove user data and token from localStorage
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    
    setUser(null);
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, user, login, signup, logout }}>
      {children}
    </AuthContext.Provider>
  );
};