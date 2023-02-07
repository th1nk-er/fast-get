import service from "./service";

export interface Result<T> {
    code: number;
    data: T;
    msg: string;
}


export const getUserMessage = (msgID: string, systemKey: string, msgKey: string, raw: boolean = false) => {
    return service.get<Result<any> | any>(`/msg`,
        {
            params:
            {
                "id": msgID == '' ? null : msgID,
                'sk': systemKey == '' ? null : systemKey,
                'mk': msgKey == '' ? null : msgKey,
                "r": raw == false ? null : raw,
            }
        })
}

export const saveUserMessage = (msg: string, msgKey: string, systemKey: string, editKey: string) => {
    return service.post<Result<any>>(`/msg`,
        {
            "m": msg == '' ? null : msg,
            'mk': msgKey == '' ? null : msgKey,
            'sk': systemKey == '' ? null : systemKey,
            'ek': editKey == '' ? null : editKey,
        })
}

export const editUserMessage = (msgID: string, msg: string, systemKey: string, editKey: string) => {
    return service.put<Result<any>>(`/msg`,
        {
            "id": msgID == '' ? null : msgID,
            "m": msg == '' ? null : msg,
            'sk': systemKey == '' ? null : systemKey,
            'ek': editKey == '' ? null : editKey,
        })
}
