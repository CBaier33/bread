import React, { useState } from "react";
import { DotsHorizontalIcon, Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { IconButton } from "@radix-ui/themes";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { models } from "../../wailsjs/go/models";

interface TransactionCardProps {
  transaction: models.Transaction;
  width?: string;
}

const TransactionCard: React.FC<TransactionCardProps> = ({ transaction, width }) => {

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
                  onClick={() => null}
                >
                  <Pencil1Icon className="w-4 h-4 text-black" /> Edit
                </div>
              </DropdownMenu.Item>

              {/* Delete item */}
              <DropdownMenu.Item asChild>
                <div
                  className="flex items-center gap-2 px-3 py-2 text-sm rounded cursor-pointer text-red-600 hover:bg-red-100"
                  onClick={() => null}
                >
                  <TrashIcon className="w-4 h-4 red" /> Delete
                </div>
              </DropdownMenu.Item>
            </DropdownMenu.Content>
          </DropdownMenu.Portal>
        </DropdownMenu.Root>

       {/* Modals rendered outside the dropdown */}
      </div>
      <div className="flex flex-col justify-between ">
        <p className="text-gray-600">$ {(transaction.amount / 100).toFixed(2)}</p>
        <p className="text-gray-600">{new Date(transaction.date).toLocaleDateString()}</p>
      </div>
    </div>
  );
};

export default TransactionCard;

