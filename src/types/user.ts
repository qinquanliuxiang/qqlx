export interface LoginRequestType {
  email: string;
  password: string;
}

export interface LoginResponseType {
  user: User;
  token: string;
}

export interface User {
  id: number;
  createdAt: number;
  updatedAt: number;
  deletedAt: number;
  name: string;
  avatar: string;
  email: string;
  mobile: string;
  roleID: number;
  role: Role;
  status: number;
}

export interface Role {
  id: number;
  createdAt: number;
  updatedAt: number;
  deletedAt: number;
  name: string;
  description: string;
  policys: Ppolicys[] | null;
}

export interface Ppolicys {
  id: number;
  createdAt: number;
  updatedAt: number;
  deletedAt: number;
  name: string;
  path: string;
  method: string;
  description: string;
}
