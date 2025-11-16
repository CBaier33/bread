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

interface BudgetProps {
  globalProject: models.Project;
  setGlobalProject: (project: models.Project) => void;
  projectList: models.Project[];
}

const Budgets: React.FC<BudgetProps> = ({ globalProject, setGlobalProject, projectList }) => {
  const [budgets, setBudgets] = useState<models.Budget[]>([]);
  const [name, setName] = useState("");

  // Fetch budgets from backend
  const fetchBudgets = async () => {
    try {
      const result = await ListBudgets(globalProject.id);
      setBudgets(result ?? []);
    } catch (err) {
      console.error("Failed to fetch budgets:", err);
    }
  };

  useEffect(() => {
    fetchBudgets();
  }, [globalProject]);


  const handleCreateBudget = async (name: string, startDate: Date | undefined, endDate: Date | undefined, expectedIncome: number, startingBalance: number) => {

    const newstart = startDate?.toISOString() ?? ""
    const newend = endDate?.toISOString() ?? ""

    console.log("newstart", newstart)
    console.log("newend", newend)

    try {
      await CreateBudget(globalProject.id, name, newstart, newend, expectedIncome, startingBalance);
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
      <ProjectSelect globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projectList}/>
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

