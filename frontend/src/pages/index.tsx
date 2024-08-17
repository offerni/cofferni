import { API_BASE_URL } from '#/common/constants';
import {
    APIItemsList,
    CreateOrderOpts,
    ItemsList,
    Order,
} from '#/common/types';
import {
    convertAPIItemsListToDomain,
    convertAPIOrderToDomain,
    convertDomainOrderOptsToAPI,
} from '#/common/utils';
import { MessageDialog } from '#/components/MessageDialog';
import Image from 'next/image';
import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'react-toastify';

export async function getServerSideProps() {
    try {
        const items: APIItemsList = await fetch(`${API_BASE_URL}/items`).then(
            (response) => {
                if (response.ok) {
                    return response.json();
                }
                throw new Error('Failed to fetch items');
            }
        );

        return {
            props: {
                items: convertAPIItemsListToDomain(items.data),
            },
        };
    } catch (error) {
        console.error(error);

        return {
            props: {
                items: [],
                error: 'Something went wrong',
            },
        };
    }
}

interface HomeProps {
    items: ItemsList['data'];
    error?: string;
}
export default function Home({ items, error }: HomeProps) {
    const { register, handleSubmit, setValue, watch, reset } =
        useForm<CreateOrderOpts>({
            defaultValues: {
                quantity: 0,
            },
        });

    const [observationFieldOpen, setObservationFieldOpen] = useState(false);

    const [createdOrder, setCreatedOrder] = useState<Order | null>(null);

    useEffect(() => {
        if (error) {
            toast(error, { type: 'error' });
        }
    }, [error]);

    const selectedItemId = watch('itemId');

    useEffect(() => {
        if (selectedItemId) {
            setValue('quantity', 1);
        }
    }, [selectedItemId, setValue]);

    const placeOrder = async (data: CreateOrderOpts) => {
        try {
            const response = await fetch(`${API_BASE_URL}/orders`, {
                body: JSON.stringify(convertDomainOrderOptsToAPI(data)),
                headers: {
                    'Content-Type': 'application/json',
                },
                method: 'POST',
            }).then(async (response) => {
                const parsedResponse = await response.json();

                if (!response.ok) {
                    throw new Error(parsedResponse.error);
                }
                return parsedResponse;
            });
            setCreatedOrder(convertAPIOrderToDomain(response));
            reset();
        } catch (error) {
            console.log(error);
            toast((error as Error).message, { type: 'error' });
        }
    };

    return (
        <>
            <MessageDialog
                open={!!createdOrder}
                onClose={() => {
                    setCreatedOrder(null);
                }}
                title="Order placed!"
            >
                {createdOrder && (
                    <>
                        <span className="text-xl">
                            {`${createdOrder.quantity}x ${createdOrder.itemName} for ${createdOrder.customerName}`}
                        </span>
                        {createdOrder.observation && (
                            <span className="text-sm text-neutral">
                                {`Observation: ${createdOrder.observation}`}
                            </span>
                        )}
                    </>
                )}
            </MessageDialog>
            <div className="flex flex-col gap-6">
                <h2 className="text-3xl">Menu</h2>

                {error ? (
                    <p>{error}</p>
                ) : (
                    <form
                        onSubmit={handleSubmit(placeOrder)}
                        className="flex flex-col gap-6"
                    >
                        <ul className="flex flex-col gap-3">
                            {items.map((item) => (
                                <li
                                    key={item.id}
                                    className="flex gap-2 justify-between items-center"
                                >
                                    <input
                                        id={`item-${item.id}`}
                                        type="radio"
                                        value={item.id}
                                        {...register('itemId')}
                                    />
                                    <label
                                        className="flex gap-2 flex-auto"
                                        htmlFor={`item-${item.id}`}
                                    >
                                        <Image
                                            src="/coffee-beans.png"
                                            alt={item.name}
                                            width={24}
                                            height={24}
                                        />
                                        {item.name}
                                    </label>
                                    {selectedItemId === item.id && (
                                        <input
                                            {...register('quantity', {
                                                setValueAs: (value) =>
                                                    Number(value),
                                            })}
                                            type="number"
                                            min="0"
                                            className="bg-neutral text-primary w-10 text-center rounded-md"
                                            defaultValue="0"
                                        />
                                    )}
                                </li>
                            ))}
                        </ul>
                        {selectedItemId && (
                            <>
                                <button
                                    type="button"
                                    className="bg-tertiary-light text-secondary p-2 rounded-md text-xs self-center"
                                    onClick={() =>
                                        setObservationFieldOpen(
                                            !observationFieldOpen
                                        )
                                    }
                                >
                                    {`${
                                        observationFieldOpen ? '-' : '+'
                                    } Observation`}
                                </button>
                                {observationFieldOpen && (
                                    <textarea
                                        {...register('observation')}
                                        className="bg-neutral-light text-secondary placeholder:text-secondary/50 p-2 rounded-md"
                                        placeholder='Observation (e.g. "decaf")'
                                    />
                                )}
                                <input
                                    type="text"
                                    placeholder="Your name"
                                    {...register('customerName', {
                                        required: true,
                                    })}
                                    className="bg-neutral-light rounded-md p-2 text-secondary placeholder:text-secondary/50"
                                />
                            </>
                        )}
                        <button className="bg-primary-light text-neutral p-2 rounded-md w-full">
                            Place order
                        </button>
                    </form>
                )}
            </div>
        </>
    );
}
