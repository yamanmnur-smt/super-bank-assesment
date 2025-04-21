import { BaseResponse } from "@/context/interfaces/base_response";

const base_url = process.env.NEXT_PUBLIC_CLIENT_URL || "http://localhost:3000";
const routes = {
  login: `${base_url}/api/login`,
}

export interface LoginRequestBody {
  username: string;
  password: string;
}

export interface UserProfileResponse {
  id : number
  name : string
  username : string
}

export interface LoginResponseBody {
  user : UserProfileResponse
  token : string
}

export const LoginRequest = async (requestBody : LoginRequestBody) : Promise<BaseResponse<LoginResponseBody>> => {
  const res = await fetch(`${routes.login}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(requestBody),
  });
  if (!res.ok) {
    throw new Error("Invalid credentials");
  }
  const data : BaseResponse<LoginResponseBody> = await res.json();
  return data;
};
