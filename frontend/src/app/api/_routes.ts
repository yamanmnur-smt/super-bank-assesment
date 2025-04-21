
const base_url = process.env.NEXT_PUBLIC_API_URL

const routes = {
    login: `${base_url}/api/v1/auth/login`,
    customer: `${base_url}/api/v1/customer/list`,
    customer_detail: `${base_url}/api/v1/customer/detail`,
    dashboard: `${base_url}/api/v1/dashboard`,
}

export default routes
  