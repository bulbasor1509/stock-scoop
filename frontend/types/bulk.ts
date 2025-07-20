export interface Bulk{
    date: string
    symbol: string
    security: string
    client: string
    type: string
    trade_qty: number,
    average_price: number
}


export interface BulkDataResponse {
    id: number
    message: string
    data: Bulk[]
}