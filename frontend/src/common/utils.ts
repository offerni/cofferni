import {
    APICreateOrderOpts,
    APIItem,
    APIOrder,
    CreateOrderOpts,
    Item,
    Order,
} from './types';

export const convertAPIItemToDomain = (item: APIItem): Item => {
    return {
        available: item.available,
        createdAt: item.created_at,
        description: item.description,
        id: item.id,
        modifiedAt: item.modified_at,
        name: item.name,
    };
};

export const convertAPIItemsListToDomain = (itemsList: APIItem[]): Item[] => {
    return itemsList.map(convertAPIItemToDomain);
};

export const convertAPIOrderToDomain = (order: APIOrder): Order => {
    return {
        createdAt: order.created_at,
        customerName: order.customer_name,
        id: order.id,
        itemId: order.item_id,
        itemName: order.item_name,
        modifiedAt: order.modified_at,
        observation: order.observation,
        quantity: order.quantity,
    };
};

export const convertDomainOrderOptsToAPI = (
    order: CreateOrderOpts
): APICreateOrderOpts => {
    return {
        customer_name: order.customerName,
        item_id: order.itemId,
        observation: order.observation,
        quantity: order.quantity,
    };
};

export const convertAPIOrdersListToDomain = (
    ordersList: APIOrder[]
): Order[] => {
    return ordersList.map(convertAPIOrderToDomain);
};
