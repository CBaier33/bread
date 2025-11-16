import { Flex } from "@radix-ui/themes";
import React, { useEffect, useState } from "react";
import { models } from "../../wailsjs/go/models";
import ProjectSelect from "../components/ProjectSelect";
import Categories from "./Categories";
import Groups from "./Groups";
import Projects from "./Projects";
import Transactions from "./Transactions";

interface HomeProps {
  globalProject: models.Project;
  setGlobalProject: (project: models.Project) => void;
  projects: models.Project[];
  globalGroup: models.Group;
  setGlobalGroup: (group: models.Group) => void;
  groupList: models.Group[];
  globalBudget: models.Budget;
  setGlobalBudget: (budget: models.Budget) => void;
  budgetList: models.Budget[];
}

const Home: React.FC<HomeProps> = (
    { 
      globalProject, setGlobalProject, projects,
      globalGroup, setGlobalGroup, groupList,
      globalBudget, setGlobalBudget, budgetList,
    }
  ) => {

  return (
    <div className="p-5 h-[calc(100vh-4rem)]">
      <div className="flex items-center gap-2 justify-center mb-4">
        <div className="flex w-full justify-start">
          <h1 className="text-4xl font-bold">*Budget Name*</h1>
        </div>
        <div className="flex w-full justify-end">
          <ProjectSelect
            globalProject={globalProject}
            setGlobalProject={setGlobalProject}
            projectList={projects}
          />
        </div>
      </div>

      <div className="w-full h-full grid grid-cols-4 gap-4">
        <div className="w-full flex flex-col col-span-3 gap-2">
          <div className="grid grid-cols-3 gap-1 ">
            <div className="flex shadow-lg rounded-xl outline outline-zinc-200 justify-center">
              <h1 className="text-3xl font-bold">*Balance*</h1>
            </div>
            <div className="flex shadow-lg rounded-xl outline outline-zinc-200 justify-center">
              <h1 className="text-3xl text-red font-bold">*Cost*</h1>
            </div>
            <div className="flex shadow-lg rounded-xl outline outline-zinc-200 justify-center">
              <h1 className="text-3xl font-bold">*Allocations*</h1>
            </div>
          </div>

          <div className="flex-1 overflow-auto shadow-lg rounded-xl outline outline-zinc-200 gap-2 mt-2 p-2">
            <Groups
              globalProject={globalProject}
              setGlobalProject={setGlobalProject}
              projectList={projects}
            />
          </div>
        </div>
        <div className="h-full overflow-auto shadow-lg rounded-xl outline outline-zinc-200 gap-2 p-0">
          <Transactions
            globalProject={globalProject}
            setGlobalProject={setGlobalProject}
            projectList={projects}
          />
        </div>
      </div>
    </div>
  );
};

export default Home;

