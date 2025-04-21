import { cookies } from "next/headers";
import routes from "../_routes";
import {
  PageResponse,
  responseError,
  responseUnauth,
} from "@/context/interfaces/base_response";
import { CustomerData } from "@/app/(main)/customers/_dto/customers_dto";

export async function GET(request: Request) {
  try {
    const cookies_bearer_token = (await cookies()).get("bearer_token");
    const token = cookies_bearer_token?.value;
    if (!token) {
      return new Response(JSON.stringify(responseUnauth), { status: 401 });
    }
    const { searchParams } = new URL(request.url);
    const page_number = searchParams.get("page_number") || "0";
    const page_size = searchParams.get("page_size") || "10";
    const search = searchParams.get("search") || "";
    const sort_by = searchParams.get("sort_by") || "";
    const sort_direction = searchParams.get("sort_direction") || "desc";

    const response = await fetch(
      `${routes.customer}?page_number=${page_number}&page_size=${page_size}&search=${search}&sort_by=${sort_by}&sort_direction=${sort_direction}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    const data: PageResponse<CustomerData> = await response.json();
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
