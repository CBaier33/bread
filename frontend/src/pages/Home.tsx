import { Flex } from "@radix-ui/themes";
import React, { useEffect, useState } from "react";
import { CreateBudget, ListBudgets } from "../../wailsjs/go/controllers/BudgetController";
import { models } from "../../wailsjs/go/models";
import BudgetSelect from "../components/BudgetSelect";
import NewBudgetPrompt from "../components/NewBudgetPrompt";
import ProjectSelect from "../components/ProjectSelect";
import { GlobalStore } from "../hooks/useGlobalStore";
import { useAppStore } from "../stores/useAppStore";
import Categories from "./Categories";
import Groups from "./Groups";
import Projects from "./Projects";
import Transactions from "./Transactions";

interface HomeProps {
  appStore: ReturnType<typeof useAppStore>;
}

const Home: React.FC<HomeProps> = ( { appStore, }) => {

  const {
    selectedProject: currentProject,
    setSelectedProject: setProject,
    projects: projects,

    budgets: budgets,
    selectedBudget: currentBudget,
    setSelectedBudget: setBudget,

    groups: groups,
  } = appStore;

  const handleCreateBudget = async (name: string, startDate: Date | undefined, endDate: Date | undefined, expectedIncome: number, startingBalance: number) => {

    const newstart = startDate?.toISOString() ?? ""
    const newend = endDate?.toISOString() ?? ""

    console.log("newstart", newstart)
    console.log("newend", newend)

    try {
      await CreateBudget(currentProject?.id ?? 0, name, newstart, newend, expectedIncome, startingBalance);
    } catch (err) {
      console.error("Failed to create budget:", err);
    }

  };


  return (
    <div className="p-5 h-[calc(100vh-4rem)]">
      <div className="flex items-center gap-2 justify-center mb-4">
        <div className="flex w-full justify-start items-center gap-2 ">
          <BudgetSelect budget={currentBudget} setBudget={setBudget} budgets={budgets} />
          <NewBudgetPrompt onSave={handleCreateBudget}/>
        </div>
        <div className="flex w-full justify-end">
          <ProjectSelect
              globalProject={currentProject}
              setGlobalProject={setProject}
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
              appStore={appStore}
            />
          </div>
        </div>
        <div className="h-full overflow-auto shadow-lg rounded-xl outline outline-zinc-200 gap-2 p-0">
          <Transactions
            appStore={appStore}
          />
        </div>
      </div>
    </div>
  );
};

export default Home;

