import React, { useEffect, useState } from "react";
import { Button, Popover, TextArea, IconButton, Dialog, Text, TextField, Box, Flex, Avatar, Checkbox } from "@radix-ui/themes";
import {PlusIcon, ChatBubbleIcon} from "@radix-ui/react-icons"
import { DayPicker } from "react-day-picker";
import "react-day-picker/dist/style.css";
import "../budget.css"
import dayjs from "dayjs";

import { ListBudgets, CreateBudget } from "../../wailsjs/go/controllers/BudgetController";
import { models } from "../../wailsjs/go/models";
import NewBudgetPrompt from "../components/NewBudgetPrompt";
import ProjectSelect from "../components/ProjectSelect";
import { useAppStore } from "../stores/useAppStore";

interface BudgetProps {
  appStore: ReturnType<typeof useAppStore>;
}

const Budgets: React.FC<BudgetProps> = ({ appStore }) => {
  const [budgets, setBudgets] = useState<models.Budget[]>([]);
  const [name, setName] = useState("");

  const {
    selectedProject: currentProject,
    setSelectedProject: setProject,
    projects: projects,

    selectedBudget: currentBudget,
  } = appStore;

  // Fetch budgets from backend
  const fetchBudgets = async () => {
    try {
      const result = await ListBudgets(currentProject?.id ?? 0);
      setBudgets(result ?? []);
    } catch (err) {
      console.error("Failed to fetch budgets:", err);
    }
  };

  useEffect(() => {
    fetchBudgets();
  }, [currentProject]);


  const handleCreateBudget = async (name: string, startDate: Date | undefined, endDate: Date | undefined, expectedIncome: number, startingBalance: number) => {

    const newstart = startDate?.toISOString() ?? ""
    const newend = endDate?.toISOString() ?? ""

    console.log("newstart", newstart)
    console.log("newend", newend)

    try {
      await CreateBudget(currentProject?.id ?? 0, name, newstart, newend, expectedIncome, startingBalance);
    } catch (err) {
      console.error("Failed to create budget:", err);
    }

    fetchBudgets();

    // Clear inputs
    setName("");

  };

  return (
    <div className="p-5"> 
      <div className="mb-4 flex gap-2 items-center w-full">
      <h1 className="text-4xl font-bold">Budgets</h1>
      <ProjectSelect appStore={appStore}/>
      <NewBudgetPrompt onSave={handleCreateBudget}/>
    </div>

      {/* Budget List */}
      <ul className="mb-6">
        {budgets.map((b) => (
          <li key={b.id} className="mb-2 p-2 border rounded">
            <strong>{b.name}</strong> (
            {dayjs(b.period_start).format("YYYY-MM-DD")} - {dayjs(b.period_end).format("YYYY-MM-DD")})
          </li>
        ))}
      </ul>
      
    </div>
  );
};

export default Budgets;

