import React, { useEffect, useState } from "react";
import { models } from "../../wailsjs/go/models";
import { ListProjects, CreateProject, UpdateProject, DeleteProject } from "../../wailsjs/go/controllers/ProjectController";
import NewProjectPrompt from "../components/NewProjectPrompt";
import ProjectCard from "../components/ProjectCard";
import { Button } from "@radix-ui/themes";

interface ProjectsProps {
  globalProject: models.Project;
  setGlobalProject: (project: models.Project) => void;
}

const Projects: React.FC<ProjectsProps> = ({ globalProject, setGlobalProject }) => {
  const [projects, setProjects] = useState<models.Project[]>([]);

  const fetchProjects = async () => {
    try {
      const result = await ListProjects();
      setProjects(result ?? []);
    } catch (err) {
      console.error("Failed to fetch projects:", err);
    }
  };

  useEffect(() => {
    fetchProjects();
  }, []);

  const handleCreateProject = async (name: string, description: string, currency: string) => {
    try {
      await CreateProject(name, description, currency);
      await fetchProjects();
    } catch (err) {
      console.error("Failed to create project:", err);
    }
  };

  const handleEditProject = async (project: models.Project) => {
    try {
      await UpdateProject(project);
      await fetchProjects();
    } catch (err) {
      console.error("Failed to create project:", err);
    }
  };

  const handleDeleteProject = async (project: models.Project) => {
    try {
      await DeleteProject(project.id);
      await fetchProjects();
    } catch (err) {
      console.error("Failed to create project:", err);
    }
  };

  return (
<div className="p-5">
  {/* Header stays left/right */}
  <div className="mb-4 flex gap-2 items-center w-full">
    <h1 className="text-4xl font-bold">Projects</h1>
    <NewProjectPrompt onSave={handleCreateProject} />
  </div>

  {/* Cards container */}
    <div className="w-full gap-7">
      {projects.map((p) => (
        <ProjectCard
          key={p.id}
          project={p}
          onEdit={handleEditProject}
          onDelete={handleDeleteProject}
          width=""
        />
      ))}
  </div>
</div>
  );
};

export default Projects;

