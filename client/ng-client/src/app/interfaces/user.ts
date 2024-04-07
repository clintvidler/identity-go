export interface User {
  id: number;
  displayName: string;
  email: string;
  created_at: string;
  updated_at: string;
  expired_at: string;
}

export interface LoginCredential {
  email: string;
  password: string;
}
