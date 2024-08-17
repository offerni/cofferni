import '#/styles/globals.css';
import type { AppProps } from 'next/app';
import Link from 'next/link';
import { ToastContainer } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

export default function App({ Component, pageProps }: AppProps) {
    return (
        <main className="flex flex-col gap-6 p-6 md:w-1/2 max-w-full mx-auto">
            <div className="flex justify-between">
                <h1 className="text-5xl font-extralight text-primary-light">
                    Cofferni
                </h1>
                <Link href="/">Menu</Link>
            </div>
            <ToastContainer />
            <Component {...pageProps} />
        </main>
    );
}
