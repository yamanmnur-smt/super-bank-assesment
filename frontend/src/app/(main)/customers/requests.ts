import { BaseResponse, PageResponse } from "@/context/interfaces/base_response";
import { CustomerData, CustomerDetailData, PageRequest } from "./_dto/customers_dto";

const base_url = process.env.NEXT_PUBLIC_CLIENT_URL || "http://localhost:3000";
const routes = {
  customer_list: `${base_url}/api/customer`,
  customer_detail: `${base_url}/api/customer`,
};

function objectToQueryParams(obj: Record<string, any>): string {
  const params = new URLSearchParams();

  Object.entries(obj).forEach(([key, value]) => {
    if (value !== undefined && value !== null) {
      params.set(key, value.toString());
    }
  });

  return params.toString();
}

export const GetCustomers = async (
  requestBody: PageRequest
): Promise<PageResponse<CustomerData>> => {

  const params = objectToQueryParams(requestBody)
  const res = await fetch(`${routes.customer_list}?${params}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const data: PageResponse<CustomerData> = await res.json()
  return data;
};

export const GetCustomerDetail = async (
  id: number
): Promise<BaseResponse<CustomerDetailData>> => {

  const res = await fetch(`${routes.customer_detail}/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const data: BaseResponse<CustomerDetailData> = await res.json()
  return data;
};