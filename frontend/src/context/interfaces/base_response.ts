interface MetaData {
    status : string;
    message : string;
    code : string
}

interface PageData {
    limit : number
    page : number
    sort : string
    total_rows : number
    total_pages : number
}

interface BaseResponse<T> {
    meta_data : MetaData;
    data : T;
}

interface MetaResponse {
    meta_data : MetaData;
}

interface PageResponse<T> {
    meta_data : MetaData;
    page_data : PageData;
    data : T[];
}

const responseError : MetaResponse = {
    meta_data : {
        status : "failed",
        code : "500",
        message : "Internal Server Error",
    }
}
const responseUnauth : MetaResponse = {
    meta_data : {
        status : "failed",
        code : "401",
        message : "Unauthorized",
    }
}
export {responseError, responseUnauth}
export type {PageResponse, MetaData, BaseResponse, MetaResponse};
