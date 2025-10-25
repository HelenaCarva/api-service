// Timestamp: 2025-10-25 01:12:56

interface ApiResponse<T> {
  data: T | null;
  error: string | null;
  statusCode: number;
}

