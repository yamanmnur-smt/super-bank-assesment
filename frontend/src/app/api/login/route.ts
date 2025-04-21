import type { NextApiRequest, NextApiResponse } from "next";
import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import {
  BaseResponse,
  responseError,
  responseUnauth,
} from "@/context/interfaces/base_response";
import { LoginRequestBody, LoginResponseBody } from "@/app/(auth)/login/requests";
import routes from "../_routes";

export async function POST(request: Request) {
  try {
    const requestBody : LoginRequestBody = await request.json();

    const response = await fetch(routes.login, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestBody),
    });

    const data: BaseResponse<LoginResponseBody> = await response.json();

    if (response.ok && data.data.token) {
      const res = NextResponse.json(data);
      (await cookies()).set({
        name: "bearer_token",
        value: data.data.token,
        httpOnly: true,
        secure: false,
        path: "/",
      });

      return new Response(
        JSON.stringify({
          meta_data: {
            status: "success",
            code: "200",
            message: "success",
          },
          data: {
            user: {
              id: data.data.user.id,
              username: data.data.user.username,
              name: data.data.user.name,
            },
          },
        }),
        { status: 200 }
      );
    }

    return new Response(JSON.stringify(responseUnauth), { status: 401 });
  } catch (error) {
    return new Response(JSON.stringify(responseError), { status: 500 });
  }
}
