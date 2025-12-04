import { createGlobalStore } from "../hooks/useGlobalStore";
import { models } from "../../wailsjs/go/models";
import { ListGroups } from "../../wailsjs/go/controllers/GroupController";

export const useGroupStore = (projectId: number) =>
  createGlobalStore<models.Group>(
    () => ListGroups(projectId),
    new models.Group({
      id: 0,
      description: "",
    })
  )();

