import { cookies } from "next/headers";
import routes from '../../_routes';


export async function GET(_request: Request, context: { params: { id: string } }) {
    try {
        const { id } = await context.params;

        const cookies_bearer_token = (await cookies()).get("bearer_token")
        const token = cookies_bearer_token?.value

        const response = await fetch(
            `${routes.customer_detail}/${id}`,
            {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization' : `Bearer ${token}`
                },
            }
        );

        const data = await response.json();
        if (response.ok && data.meta_data.status === 'success') {
            return new Response(
                JSON.stringify({
                    meta_data: data.meta_data,
                    data: data.data,
                    page_data: data.page_data,
                }),
                { status: 200 }
            );
        }

        return new Response(
            JSON.stringify({
                meta_data: {
                    status: 'failed',
                    code: '401',
                    message: 'Invalid Credential',
                },
            }),
            { status: 401 }
        );
       
    } catch (error) {
        return new Response(
            JSON.stringify({
                meta_data: {
                    status: 'failed',
                    code: '500',
                    message: 'Internal Server Error',
                },
                error : error
            }),
            { status: 500 }
        );
    }
}
