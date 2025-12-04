import React, { useEffect, useState } from "react";
import { models } from "../../wailsjs/go/models";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { CheckIcon, ChevronDownIcon } from "@radix-ui/react-icons";
import { GlobalStore } from "../hooks/useGlobalStore";
import { ListBudgets } from "../../wailsjs/go/controllers/BudgetController";

interface BudgetSelectProps {
  budget: models.Budget | null;
  budgets: models.Budget[];
  setBudget: (budget: models.Budget) => void; 
}

const BudgetSelect: React.FC<BudgetSelectProps> = ({
  budget,
  budgets,
  setBudget,
}) => {

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger asChild>
        <button className="inline-flex items-center gap-1 rounded-lg border border-gray-300 bg-white px-3 py-1.5 text-sm shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500">
          {budget?.name}
          <ChevronDownIcon className="h-4 w-4 opacity-70" />
        </button>
      </DropdownMenu.Trigger>

      <DropdownMenu.Portal>
        <DropdownMenu.Content
          sideOffset={4}
          className="min-w-[160px] rounded-lg border border-gray-200 bg-white p-1 shadow-lg"
        >
          <DropdownMenu.RadioGroup
            value={budget?.name}
            onValueChange={(name) => {
              const proj = budgets.find((p) => p.name === name);
              if (proj) setBudget(proj);
            }}
            className="flex flex-col gap-0.5"
          >
            {budgets.map((b) => (
              <DropdownMenu.RadioItem
                key={b.name}
                value={b.name}
                className="budget flex cursor-pointer select-none items-center gap-2 rounded-md px-2 py-1.5 text-sm text-gray-700 hover:bg-gray-100 focus:bg-gray-100 focus:outline-none"
              >
                <DropdownMenu.ItemIndicator>
                  <CheckIcon className="h-4 w-4 text-blue-500" />
                </DropdownMenu.ItemIndicator>
                <span className="capitalize">{b.name}</span>
              </DropdownMenu.RadioItem>
            ))}
          </DropdownMenu.RadioGroup>
        </DropdownMenu.Content>
      </DropdownMenu.Portal>
    </DropdownMenu.Root>
  );
};

export default BudgetSelect;

