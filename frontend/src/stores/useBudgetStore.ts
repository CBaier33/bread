import { createGlobalStore } from "../hooks/useGlobalStore";
import { models } from "../../wailsjs/go/models";
import { ListBudgets } from "../../wailsjs/go/controllers/BudgetController";

export const useBudgetStore = (projectId: number | null) =>
  createGlobalStore<models.Budget>(
    () => ListBudgets(projectId ?? 0),
    new models.Budget({
      id: 0,
      description: "",
      name: "Select Budget"
    })
  )();

