import { useState, useCallback } from 'react';
import { User } from '../types';
import { api } from '../services/api';

export function useAuth() {
  const [user, setUser] = useState<User | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const login = useCallback(async (email: string, password: string) => {
    try {
      const { data, error } = await api.auth.login(email, password);
      if (error) throw error;
      if (data) {
        setUser(data);
        setIsAuthenticated(true);
      }
    } catch (error) {
      console.error('Login error:', error);
      throw error;
    }
  }, []);

  const logout = useCallback(async () => {
    try {
      const { error } = await api.auth.logout();
      if (error) throw error;
      setUser(null);
      setIsAuthenticated(false);
    } catch (error) {
      console.error('Logout error:', error);
      throw error;
    }
  }, []);

  return {
    user,
    isAuthenticated,
    login,
    logout,
  };
}