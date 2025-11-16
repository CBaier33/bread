import React from "react";
import * as AlertDialog from "@radix-ui/react-alert-dialog";

interface DeleteTransactionDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onConfirm: () => Promise<void> | void;
}

const DeleteTransactionDialog: React.FC<DeleteTransactionDialogProps> = ({ open, onOpenChange, onConfirm }) => {
  const handleConfirm = async () => {
    await onConfirm();
    onOpenChange(false);
  };

  return (
    <AlertDialog.Root open={open} onOpenChange={onOpenChange}>
      		<AlertDialog.Portal>
			<AlertDialog.Overlay className="fixed inset-0 bg-blackA6 data-[state=open]:animate-overlayShow" />
			<AlertDialog.Content className="fixed left-1/2 top-1/2 max-h-[85vh] w-[90vw] max-w-[500px] -translate-x-1/2 -translate-y-1/2 rounded-md bg-gray1 p-[25px] shadow-[var(--shadow-6)] focus:outline-none data-[state=open]:animate-contentShow">
				<AlertDialog.Title className="m-0 text-[17px] font-medium text-mauve12">
					Are you absolutely sure?
				</AlertDialog.Title>
				<AlertDialog.Description className="mb-5 mt-[15px] text-[15px] leading-normal text-mauve11">
					This action cannot be undone. This will permanently delete your
					transaction.
				</AlertDialog.Description>
				<div className="flex justify-end gap-[25px]">
					<AlertDialog.Cancel asChild>
						<button className="inline-flex h-[35px] items-center justify-center rounded bg-mauve4 px-[15px] font-medium leading-none text-mauve11 outline-none outline-offset-1 hover:bg-mauve5 focus-visible:outline-2 focus-visible:outline-mauve7 select-none">
							Cancel
						</button>
					</AlertDialog.Cancel>
					<AlertDialog.Action asChild>
						<button className="inline-flex h-[35px] items-center justify-center rounded bg-red4 px-[15px] font-medium leading-none text-red11 outline-none outline-offset-1 hover:bg-red5 focus-visible:outline-2 focus-visible:outline-red7 select-none" onClick={handleConfirm}>
							Yes, delete transaction
						</button>
					</AlertDialog.Action>
				</div>
			</AlertDialog.Content>
		</AlertDialog.Portal>
    </AlertDialog.Root>
  );
};

export default DeleteTransactionDialog;

