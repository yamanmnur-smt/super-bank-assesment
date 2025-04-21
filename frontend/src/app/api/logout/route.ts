import {cookies} from "next/headers"

export async function POST(request: Request) {
    (await cookies()).delete("bearer_token")
    return new Response(JSON.stringify(""), { status: 200 });
}
