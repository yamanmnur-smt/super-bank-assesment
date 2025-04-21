import { cookies } from "next/headers";
import routes from "../_routes";
import {
    BaseResponse,
  responseError,
  responseUnauth,
} from "@/context/interfaces/base_response";
import { DashboardData } from "@/app/(main)/dashboard/_dto/dashboard_dto";

export async function GET(request: Request) {
  try {
    const cookies_bearer_token = (await cookies()).get("bearer_token");
    const token = cookies_bearer_token?.value;
    if (!token) {
      return new Response(JSON.stringify(responseUnauth), { status: 401 });
    }
  
    const response = await fetch(
      `${routes.dashboard}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    const data: BaseResponse<DashboardData> = await response.json();
    if (response.ok && data.meta_data.status === "success") {
      return new Response(JSON.stringify(data), { status: 200 });
    }

    (await cookies()).delete("bearer_token");

    return new Response(JSON.stringify(responseUnauth), { status: 401 });
  } catch (error) {
    (await cookies()).delete("bearer_token");

    return new Response(JSON.stringify(responseError), { status: 500 });
  }
}
