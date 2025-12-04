import { createGlobalStore } from "../hooks/useGlobalStore";
import { models } from "../../wailsjs/go/models";
import { ListProjects } from "../../wailsjs/go/controllers/ProjectController";

export const useProjectStore = createGlobalStore<models.Project>(
  () => ListProjects(),
  new models.Project({
    id: 0,
    name: "Empty",
    description: "",
    currency: "",
  })
);

