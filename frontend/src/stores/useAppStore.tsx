import { useEffect, useState } from "react";
import { models } from "../../wailsjs/go/models";

import { ListProjects } from "../../wailsjs/go/controllers/ProjectController";
import { ListBudgets } from "../../wailsjs/go/controllers/BudgetController";
import { ListGroups } from "../../wailsjs/go/controllers/GroupController";

export function useAppStore() {
  const [projects, setProjects] = useState<models.Project[]>([]);
  const [budgets, setBudgets] = useState<models.Budget[]>([]);
  const [groups, setGroups] = useState<models.Group[]>([]);

  const [selectedProject, setSelectedProject] = useState<models.Project | null>(null);
  const [selectedBudget, setSelectedBudget] = useState<models.Budget | null>(null);
  const [selectedGroup, setSelectedGroup] = useState<models.Group | null>(null);

  const refreshProjects = async () => {
    const result = await ListProjects();
    const safe = result ?? [];

    setProjects(safe);

    if (safe.length > 0) {
      setSelectedProject((prev) => prev ?? safe[0]);
    } else {
      setSelectedProject(null);
    }
  };

  const refreshBudgets = async (projectID: number) => {
    const result = await ListBudgets(projectID);
    const safe = result ?? [];

    setBudgets(safe);

    if (safe.length > 0) {
      setSelectedBudget((prev) => prev ?? safe[0]);
    } else {
      setSelectedBudget(null);
    }
  };

  const refreshGroups = async (projectID: number) => {
    const result = await ListGroups(projectID);
    const safe = result ?? [];

    setGroups(safe);

    if (safe.length > 0) {
      setSelectedGroup((prev) => prev ?? safe[0]);
    } else {
      setSelectedGroup(null);
    }
  };

  useEffect(() => {
    refreshProjects();
  }, []);

  useEffect(() => {
    if (!selectedProject) return;

    refreshBudgets(selectedProject.id);
    refreshGroups(selectedProject.id);
  }, [selectedProject]);

  return {
    // data
    projects,
    budgets,
    groups,

    // selections
    selectedProject,
    selectedBudget,
    selectedGroup,

    // setters
    setSelectedProject,
    setSelectedBudget,
    setSelectedGroup,

    // refreshers
    refreshProjects,
    refreshBudgets,
    refreshGroups,
  };
}

