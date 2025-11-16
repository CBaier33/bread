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

interface GroupProps {
  globalProject: models.Project;
  setGlobalProject: (project: models.Project) => void;
  projectList: models.Project[];
}

const Groups: React.FC<GroupProps> = ({ globalProject, setGlobalProject, projectList }) => {
  const [groups, setGroups] = useState<models.Group[]>([]);

  // Fetch groups from backend
  const fetchGroups = async () => {
    try {
      const result = await ListGroups(globalProject.id);
      setGroups(result ?? []);
    } catch (err) {
      console.error("Failed to fetch groups:", err);
    }
  };

  useEffect(() => {
    fetchGroups();
  }, [globalProject]);


  const handleCreateGroup = async (name: string, description: string) => {


    try {
      await CreateGroup(globalProject.id, name, description);
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

