import { API_BASE_URL } from '#/common/constants';
import { APIOrdersList, OrdersList } from '#/common/types';
import {
    convertAPIOrdersListToDomain,
    convertAPIOrderToDomain,
} from '#/common/utils';
import { useEffect, useState } from 'react';
import { toast } from 'react-toastify';

export async function getServerSideProps() {
    try {
        const pendingOrders: APIOrdersList = await fetch(
            `${API_BASE_URL}/orders?fulfilled=false`
        ).then((response) => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Failed to fetch orders');
        });

        const fulfilledOrders: APIOrdersList = await fetch(
            `${API_BASE_URL}/orders?fulfilled=true`
        ).then((response) => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Failed to fetch orders');
        });
        return {
            props: {
                pendingOrders: convertAPIOrdersListToDomain(pendingOrders.data),
                fulfilledOrders: convertAPIOrdersListToDomain(
                    fulfilledOrders.data
                ),
            },
        };
    } catch (error) {
        console.error(error);

        return {
            props: {
                orders: [],
                error: 'Something went wrong',
            },
        };
    }
}

interface OrdersProps {
    pendingOrders: OrdersList['data'];
    fulfilledOrders: OrdersList['data'];
    error?: string;
}

export default function Orders({
    pendingOrders,
    fulfilledOrders,
    error,
}: OrdersProps) {
    const [localPendingOrders, setLocalPendingOrders] = useState(pendingOrders);
    const [localFulfilledOrders, setLocalFulfilledOrders] =
        useState(fulfilledOrders);

    const [showHistory, setShowHistory] = useState(false);

    useEffect(() => {
        toast.error(error);
    }, [error]);

    const markAsDone = async (orderId: string) => {
        try {
            await fetch(`${API_BASE_URL}/orders`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    id: orderId,
                    fulfilled: true,
                }),
            }).then(async (response) => {
                if (response.ok) {
                    setLocalPendingOrders((prevOrders) =>
                        prevOrders.filter((order) => order.id !== orderId)
                    );
                    const parsedResponse = await response.json();
                    const updatedOrder =
                        convertAPIOrderToDomain(parsedResponse);
                    setLocalFulfilledOrders((prevOrders) => [
                        ...prevOrders,
                        updatedOrder,
                    ]);
                } else {
                    throw new Error('Failed to mark order as done');
                }
            });
        } catch (error) {
            toast.error('Failed to mark order as done');
        }
    };
    return (
        <div className="flex flex-col gap-6">
            <div className="flex justify-between">
                <h2 className="text-3xl">
                    {showHistory ? 'Fulfilled orders' : 'Pending orders'}
                </h2>
                <button
                    type="button"
                    className="underline text-primary-light"
                    onClick={() => setShowHistory(!showHistory)}
                >
                    {!showHistory ? 'See history' : 'Hide history'}
                </button>
            </div>
            {showHistory ? (
                <>
                    {localFulfilledOrders.length === 0 && (
                        <span className="animate-bounce text-lg">
                            No orders here!
                        </span>
                    )}
                    <ul className="grid grid-cols-1 gap-8">
                        {localFulfilledOrders.map((order) => (
                            <li key={order.id} className="flex justify-between">
                                <div className="flex flex-col gap-2">
                                    <p className="text-lg flex gap-2">
                                        <span>✓</span>
                                        <span className="text-primary-light font-bold">
                                            {order.quantity}x
                                        </span>
                                        <span className="font-bold">
                                            {order.itemName}
                                        </span>
                                        <span className="font-light capitalize">
                                            - {order.customerName}
                                        </span>
                                    </p>

                                    {order.observation && (
                                        <p className="text-sm text-primary">
                                            Observation: {order.observation}
                                        </p>
                                    )}
                                </div>
                            </li>
                        ))}
                    </ul>
                </>
            ) : (
                <>
                    {localPendingOrders.length === 0 && (
                        <span className="animate-bounce text-lg">
                            No orders here!
                        </span>
                    )}
                    <ul className="grid grid-cols-1 gap-8">
                        {localPendingOrders.map((order) => (
                            <li key={order.id} className="flex justify-between">
                                <div className="flex flex-col gap-2">
                                    <p className="text-xl flex gap-2">
                                        <span className="text-primary-light font-bold">
                                            {order.quantity}x
                                        </span>
                                        <span className="font-bold">
                                            {order.itemName}
                                        </span>
                                        <span className="font-light capitalize">
                                            - {order.customerName}
                                        </span>
                                    </p>

                                    {order.observation && (
                                        <p className="text-sm text-primary">
                                            Observation: {order.observation}
                                        </p>
                                    )}
                                </div>
                                <button
                                    onClick={() => markAsDone(order.id)}
                                    type="button"
                                    className="bg-primary-light rounded-md p-2 self-center"
                                >
                                    Done
                                </button>
                            </li>
                        ))}
                    </ul>
                </>
            )}
        </div>
    );
}
