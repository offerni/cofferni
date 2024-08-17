import { ReactNode } from 'react';
import {
    Dialog,
    DialogBackdrop,
    DialogPanel,
    DialogTitle,
} from '@headlessui/react';

interface MessageDialogProps {
    open: boolean;
    onClose: () => void;
    children: ReactNode;
    title?: string;
}

export function MessageDialog({
    open,
    onClose,
    children,
    title,
}: MessageDialogProps) {
    return (
        <Dialog
            open={open}
            as="div"
            className="relative z-10 focus:outline-none"
            onClose={onClose}
        >
            <DialogBackdrop className="fixed inset-0 bg-background/90" />
            <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
                <div className="flex min-h-full items-center justify-center p-4">
                    <DialogPanel
                        transition
                        className="w-full max-w-md rounded-xl bg-white/5 p-6 backdrop-blur-2xl duration-300 ease-out data-[closed]:transform-[scale(95%)] data-[closed]:opacity-0 min-h-56 flex flex-col gap-6"
                    >
                        {title && (
                            <DialogTitle className="text-2xl font-bold">
                                {title}
                            </DialogTitle>
                        )}
                        {children}
                    </DialogPanel>
                </div>
            </div>
        </Dialog>
    );
}
