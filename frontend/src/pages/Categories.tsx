import React, { useEffect, useState } from "react";
import { Button, Popover, TextArea, IconButton, Dialog, Text, TextField, Box, Flex, Avatar, Checkbox } from "@radix-ui/themes";
import {PlusIcon, ChatBubbleIcon} from "@radix-ui/react-icons"
import { DayPicker } from "react-day-picker";
import "react-day-picker/dist/style.css";
import "../budget.css"
import dayjs from "dayjs";

import { ListCategories, CreateCategory } from "../../wailsjs/go/controllers/CategoryController";
import { models } from "../../wailsjs/go/models";
import NewCategoryPrompt from "../components/NewCategoryPrompt";
import GroupSelect from "../components/GroupSelect";

interface CategoryProps {
  globalGroup: models.Group;
  setGlobalGroup: (group: models.Group) => void;
  groupList: models.Group[];
  globalBudget: models.Budget;
  setGlobalBudget: (budget: models.Budget) => void;
  budgetList: models.Budget[];
}

const Categories: React.FC<CategoryProps> = ({ globalBudget, setGlobalBudget, budgetList, globalGroup, setGlobalGroup, groupList }) => {
  const [budgets, setCategories] = useState<models.Category[]>([]);
  const [name, setName] = useState("");

  // Fetch budgets from backend
  const fetchCategories = async () => {
    try {
      const result = await ListCategories(globalGroup.id);
      setCategories(result ?? []);
    } catch (err) {
      console.error("Failed to fetch budgets:", err);
    }
  };

  useEffect(() => {
    fetchCategories();
  }, [globalGroup]);


  const handleCreateCategory = async (name: string, description: string,  expenseType: boolean) => {

    try {
      await CreateCategory(globalGroup.id, name, description, expenseType);
    } catch (err) {
      console.error("Failed to create budget:", err);
    }

    fetchCategories();

    // Clear inputs
    setName("");

  };

  return (
    <div className="p-5"> 
      <div className="mb-4 flex gap-2 items-center w-full">
      <h1 className="text-4xl font-bold">Categories</h1>
      <GroupSelect globalGroup={globalGroup} setGlobalGroup={setGlobalGroup} groupList={groupList}/>
      <NewCategoryPrompt onSave={handleCreateCategory} globalBudget={globalBudget} setGlobalBudget={setGlobalBudget} budgetList={budgetList}/>
    </div>



      {/* Category List */}
      <ul className="mb-6">
        {budgets.map((b) => (
          <li key={b.id} className="mb-2 p-2 border rounded">
            <strong>{b.name}</strong> - {b.description}
          </li>
        ))}
      </ul>
      
    </div>
  );
};

export default Categories;

