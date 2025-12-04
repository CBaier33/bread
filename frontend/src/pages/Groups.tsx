import React, { useEffect, useState } from "react";
import { Button, Popover, TextArea, IconButton, Dialog, Text, TextField, Box, Flex, Avatar, Checkbox } from "@radix-ui/themes";
import {PlusIcon, ChatBubbleIcon} from "@radix-ui/react-icons"
import { DayPicker } from "react-day-picker";
import "react-day-picker/dist/style.css";
import "../budget.css"
import dayjs from "dayjs";

import { ListGroups, CreateGroup } from "../../wailsjs/go/controllers/GroupController";
import { models } from "../../wailsjs/go/models";
import NewGroupPrompt from "../components/NewGroupPrompt";
import ProjectSelect from "../components/ProjectSelect";
import { useAppStore } from "../stores/useAppStore";

interface GroupProps {
  appStore: ReturnType<typeof useAppStore>;
}

const Groups: React.FC<GroupProps> = ({ appStore }) => {

  const {
    selectedProject: currentProject,
    setSelectedProject: setProject,
    projects: projects,

    budgets: budgets,
    selectedBudget: currentBudget,
    setSelectedBudget: setBudget,

  } = appStore;

  const [groups, setGroups] = useState<models.Group[]>([]);

  // Fetch groups from backend
  const fetchGroups = async () => {
    try {
      const result = await ListGroups(currentProject?.id ?? 0);
      setGroups(result ?? []);
    } catch (err) {
      console.error("Failed to fetch groups:", err);
    }
  };

  useEffect(() => {
    fetchGroups();
  }, [currentProject]);


  const handleCreateGroup = async (name: string, description: string) => {


    try {
      await CreateGroup(currentProject?.id ?? 0, name, description);
    } catch (err) {
      console.error("Failed to create group:", err);
    }

    fetchGroups();
  };

  return (
    <div className="p-5"> 
      <div className="mb-4 flex gap-2 items-center w-full">
      <h1 className="text-4xl font-bold">Groups</h1>
      <NewGroupPrompt onSave={handleCreateGroup}/>
    </div>

      {/* Group List */}
      <ul className="mb-6">
        {groups.map((b) => (
          <li key={b.id} className="mb-2 p-2 border rounded">
            <strong>{b.name}</strong> - {b.description}
          </li>
        ))}
      </ul>
      
    </div>
  );
};

export default Groups;

