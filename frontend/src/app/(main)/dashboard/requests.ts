import { BaseResponse, PageResponse } from "@/context/interfaces/base_response";
import { DashboardData } from "./_dto/dashboard_dto";

const base_url = process.env.NEXT_PUBLIC_CLIENT_URL || "http://localhost:3000";
const routes = {
  dashboard: `${base_url}/api/dashboard`,
};

export const GetDashboard = async (): Promise<BaseResponse<DashboardData>> => {

  const res = await fetch(`${routes.dashboard}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const data: BaseResponse<DashboardData> = await res.json()
  return data;
};