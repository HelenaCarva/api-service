// types.ts
export interface ErrorResponse {
  statusCode: number;
  error: string;
  message: string;
}

export interface ApiRequest {
  method: string;
  url: string;
  headers: { [key: string]: string };
  body?: any;
}

export interface ApiResponse {
  status: number;
  data: any;
  headers: { [key: string]: string };
}

export interface ServiceConfig {
  baseUrl: string;
  timeout: number;
  maxRetries: number;
}

export enum HttpMethod {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
  PATCH = 'PATCH',
}

export interface RequestOptions {
  method: HttpMethod;
  url: string;
  params?: { [key: string]: string };
  headers?: { [key: string]: string };
  body?: any;
}