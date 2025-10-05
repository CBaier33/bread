import React from "react";
import { models } from "../../wailsjs/go/models";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { CheckIcon, ChevronDownIcon } from "@radix-ui/react-icons";

interface ProjectSelectProps {
  globalProject: models.Project;
  projectList: models.Project[];
  setGlobalProject: (project: models.Project) => void;
}

const ProjectSelect: React.FC<ProjectSelectProps> = ({
  globalProject,
  projectList,
  setGlobalProject,
}) => {
  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger asChild>
        <button className="inline-flex items-center gap-1 rounded-lg border border-gray-300 bg-white px-3 py-1.5 text-sm shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500">
          {globalProject?.name ?? "Select a project"}
          <ChevronDownIcon className="h-4 w-4 opacity-70" />
        </button>
      </DropdownMenu.Trigger>

      <DropdownMenu.Portal>
        <DropdownMenu.Content
          sideOffset={4}
          className="min-w-[160px] rounded-lg border border-gray-200 bg-white p-1 shadow-lg"
        >
          <DropdownMenu.RadioGroup
            value={globalProject?.name}
            onValueChange={(name) => {
              const proj = projectList.find((p) => p.name === name);
              if (proj) setGlobalProject(proj);
            }}
            className="flex flex-col gap-0.5"
          >
            {projectList.map((project) => (
              <DropdownMenu.RadioItem
                key={project.name}
                value={project.name}
                className="group flex cursor-pointer select-none items-center gap-2 rounded-md px-2 py-1.5 text-sm text-gray-700 hover:bg-gray-100 focus:bg-gray-100 focus:outline-none"
              >
                <DropdownMenu.ItemIndicator>
                  <CheckIcon className="h-4 w-4 text-blue-500" />
                </DropdownMenu.ItemIndicator>
                <span className="capitalize">{project.name}</span>
              </DropdownMenu.RadioItem>
            ))}
          </DropdownMenu.RadioGroup>
        </DropdownMenu.Content>
      </DropdownMenu.Portal>
    </DropdownMenu.Root>
  );
};

export default ProjectSelect;

