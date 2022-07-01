export interface IUser {
    id: number
    account: string
    type: number
    type2: string
    created_at: number
    login_at: number
    remark: string

    locked: boolean
}

export interface IUserLoginLog {
    id: number
    account: string
    type: number
    time: number
    status: number
    ip: string
    address: string
}


export interface IUserSearch {
    account: string
}
