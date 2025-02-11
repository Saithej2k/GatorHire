import { useState, useCallback } from 'react';

export function useLoading(defaultState = false) {
  const [isLoading, setIsLoading] = useState(defaultState);

  const withLoading = useCallback(async <T>(fn: () => Promise<T>): Promise<T> => {
    setIsLoading(true);
    try {
      return await fn();
    } finally {
      setIsLoading(false);
    }
  }, []);

  return { isLoading, withLoading };
}