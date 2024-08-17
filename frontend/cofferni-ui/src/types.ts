// API Types
export interface APIItem {
    available: boolean
    created_at: string
    description?: string
    id: string
    modified_at: string
    name: string
}

export interface APIItemsList {
    data: APIItem[]
}

export interface APIOrder {
    created_at: string
    customer_name: string
    id: string
    item_id: APIItem['id']
    item_name: string
    modified_at: string
    observation?: string
    quantity: number
}

export interface APIOrdersList {
    data: APIOrder[]
}

export interface APICreateOrderOpts {
    customer_name: string
    item_id: APIItem['id']
    observation?: string
    quantity: number
}

export interface APIUpdateOrderOpts {
    fulfilled?: boolean
    id: APIOrder['id']
    observation?: string
    quantity?: number
}

// Domain Types
export interface Item {
    available: boolean
    createdAt: string
    description?: string
    id: string
    modifiedAt: string
    name: string
}

export interface ItemsList {
    data: Item[]
}

export interface Order {
    createdAt: string
    customerName: string
    id: string
    itemId: Item['id']
    itemName: string
    modifiedAt: string
    observation?: string
    quantity: number
}

export interface OrdersList {
    data: Order[]
}

export interface CreateOrderOpts {
    customerName: string
    itemId: Item['id']
    observation?: string
    quantity: number
}

export interface UpdateOrderOpts {
    fulfilled?: boolean
    id: Order['id']
    observation?: string
    quantity?: number
}
