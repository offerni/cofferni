// API Types
interface APIItem {
    available: boolean
    created_at: string
    description?: string
    id: string
    modified_at: string
    name: string
}

interface APIItemsList {
    data: APIItem[]
}

interface APIOrder {
    created_at: string
    customer_name: string
    id: string
    item_id: APIItem['id']
    item_name: string
    modified_at: string
    observation?: string
    quantity: number
}

interface APIOrdersList {
    data: APIOrder[]
}

interface APICreateOrderOpts {
    customer_name: string
    item_id: APIItem['id']
    observation?: string
    quantity: number
}

interface APIUpdateOrderOpts {
    fulfilled?: boolean
    id: APIOrder['id']
    observation?: string
    quantity?: number
}

// Domain Types
interface Item {
    available: boolean
    createdAt: string
    description?: string
    id: string
    modifiedAt: string
    name: string
}

interface ItemsList {
    data: Item[]
}

interface Order {
    createdAt: string
    customerName: string
    id: string
    itemId: Item['id']
    itemName: string
    modifiedAt: string
    observation?: string
    quantity: number
}

interface OrdersList {
    data: Order[]
}

interface CreateOrderOpts {
    customerName: string
    itemId: Item['id']
    observation?: string
    quantity: number
}

interface UpdateOrderOpts {
    fulfilled?: boolean
    id: Order['id']
    observation?: string
    quantity?: number
}
