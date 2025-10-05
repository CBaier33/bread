import React, { useState } from "react";
import { DotsHorizontalIcon, Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { IconButton } from "@radix-ui/themes";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { models } from "../../wailsjs/go/models";
import EditProject from "./EditProject";
import DeleteProjectDialog from "./DeleteProjectDialog";

interface ProjectCardProps {
  project: models.Project;
  onEdit: (updatedProject: models.Project) => Promise<void> | void;
  onDelete: (deleteProject: models.Project) => Promise<void> | void;
  width?: string;
}

const ProjectCard: React.FC<ProjectCardProps> = ({ project, onEdit, onDelete, width }) => {
  const [isEditOpen, setIsEditOpen] = useState(false);
  const [isDeleteOpen, setIsDeleteOpen] = useState(false);

  return (
    <div className={`p-6 mb-6 ${width} flex flex-col rounded-2xl shadow-sm border bg-white`}>
      <div className="flex justify-between items-start">
        <div>
          <h2 className="text-xl font-bold">{project.name}</h2>
          <p className="text-gray-600">{project.description}</p>
          <p>
            <span className="font-medium">Currency:</span> {project.currency}
          </p>
          <p className="text-sm text-gray-400">
            Created: {new Date(project.created_at).toLocaleDateString()}
          </p>
        </div>

        <DropdownMenu.Root>
          <DropdownMenu.Trigger asChild>
            <IconButton variant="ghost" radius="full" size="2">
              <DotsHorizontalIcon />
            </IconButton>
          </DropdownMenu.Trigger>

          <DropdownMenu.Portal>
            <DropdownMenu.Content
              align="end"
              sideOffset={5}
              className="min-w-[160px] rounded-md shadow-lg bg-white border z-50"
            >
              {/* Edit item */}
              <DropdownMenu.Item asChild>
                <div
                  className="flex items-center gap-2 px-3 py-2 text-sm rounded cursor-pointer text-black hover:bg-gray-100"
                  onClick={() => setIsEditOpen(true)}
                >
                  <Pencil1Icon className="w-4 h-4 text-black" /> Edit
                </div>
              </DropdownMenu.Item>

              {/* Delete item */}
              <DropdownMenu.Item asChild>
                <div
                  className="flex items-center gap-2 px-3 py-2 text-sm rounded cursor-pointer text-red-600 hover:bg-red-100"
                  onClick={() => setIsDeleteOpen(true)}
                >
                  <TrashIcon className="w-4 h-4 red" /> Delete
                </div>
              </DropdownMenu.Item>
            </DropdownMenu.Content>
          </DropdownMenu.Portal>
        </DropdownMenu.Root>

        {/* Modals rendered outside the dropdown */}
        <EditProject
          open={isEditOpen}
          onOpenChange={setIsEditOpen}
          project={project}
          onSave={onEdit}
        />

        <DeleteProjectDialog
          open={isDeleteOpen}
          onOpenChange={setIsDeleteOpen}
          onConfirm={() => onDelete(project)}
        />
      </div>
    </div>
  );
};

export default ProjectCard;

