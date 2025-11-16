import React, { useState } from "react";
import { DotsHorizontalIcon, Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { IconButton } from "@radix-ui/themes";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { models } from "../../wailsjs/go/models";
import DeleteTransactionDialog from "./DeleteTransactionDialog";
import EditTransaction from "./EditTransaction";

interface TransactionCardProps {
  transaction: models.Transaction;
  onEdit: (updatedTransaction: models.Transaction) => Promise<void> | void;
  onDelete: (deleteTransaction: models.Transaction) => Promise<void> | void;
  width?: string;
  categoryList: models.Category[];
}

const TransactionCard: React.FC<TransactionCardProps> = ({ transaction, width, onEdit, onDelete, categoryList}) => {
  const [isEditOpen, setIsEditOpen] = useState(false);
  const [isDeleteOpen, setIsDeleteOpen] = useState(false);

  return (
    <div className={`p-6 mb-6 ${width} justify-center rounded-2xl shadow-sm border bg-white`}>
      <div className="flex w-full justify-between">
        <p className="text-gray-600 font-bold">{transaction.description}</p>

        <DropdownMenu.Root>
          <DropdownMenu.Trigger asChild>
            <IconButton variant="ghost" radius="full" size="2">
              <DotsHorizontalIcon />
            </IconButton>
          </DropdownMenu.Trigger>

          <DropdownMenu.Portal>
            <DropdownMenu.Content
              align="end"
              sideOffset={5}
              className="min-w-[160px] rounded-md shadow-lg bg-white border z-50"
            >
              {/* Edit item */}
              <DropdownMenu.Item asChild>
                <div
                  className="flex items-center gap-2 px-3 py-2 text-sm rounded cursor-pointer text-black hover:bg-gray-100"
                  onClick={() => setIsEditOpen(true)}
                >
                  <Pencil1Icon className="w-4 h-4 text-black" /> Edit
                </div>
              </DropdownMenu.Item>

              {/* Delete item */}
              <DropdownMenu.Item asChild>
                <div
                  className="flex items-center gap-2 px-3 py-2 text-sm rounded cursor-pointer text-red-600 hover:bg-red-100"
                  onClick={() => setIsDeleteOpen(true)}
                >
                  <TrashIcon className="w-4 h-4 red" /> Delete
                </div>
              </DropdownMenu.Item>
            </DropdownMenu.Content>
          </DropdownMenu.Portal>
        </DropdownMenu.Root>

        <EditTransaction
        open={isEditOpen}
        onOpenChange={setIsEditOpen}
        transaction={transaction}
        onSave={onEdit}
        categoryList={categoryList}
      />

      <DeleteTransactionDialog
        open={isDeleteOpen}
        onOpenChange={setIsDeleteOpen}
        onConfirm={() => onDelete(transaction)}
      />

       {/* Modals rendered outside the dropdown */}
      </div>
      <div className="flex flex-row gap-1 flex-col justify-start ">
        <div
          className={`
            inline-flex items-center justify-center px-3 py-1 rounded-full text-sm font-semibold
            ${transaction.expense_type === true
              ? "text-red-600 bg-red-100"
              : "text-green-600 bg-green-100"
                }
              `}
            >
          ${(transaction.amount / 100).toFixed(2)}
        </div>
        <div className="flex flex-row justify-center gap-2">
          <div
            className="
              inline-flex w-full items-center justify-center px-3 py-1 rounded-full text-[10px] font-semibold
              bg-violet-100 text-violet-700
            "
            >
            {transaction.category_name}
          </div>
          <div
            className="
              inline-flex items-center justify-center px-3 py-1 rounded-full text-[10px] font-semibold
              bg-blue-100 text-blue-700
            "
            >
            {transaction.date}
          </div>
        </div>
      </div>
    </div>
  );
};

export default TransactionCard;

