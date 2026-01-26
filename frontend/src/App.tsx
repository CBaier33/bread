import React, { useState, useEffect } from "react";
import { Theme } from "@radix-ui/themes"
import { models } from "../wailsjs/go/models";
import { ListProjects } from "../wailsjs/go/controllers/ProjectController";
import Transactions from "./pages/Transactions";
import Projects from "./pages/Projects"
import Budgets from "./pages/Budgets"
import Analysis from "./pages/Analysis";
import Groups from "./pages/Groups";
import Categories from "./pages/Categories";
import Home from "./pages/Home";
import { ListGroups } from "../wailsjs/go/controllers/GroupController";
import { ListBudgets } from "../wailsjs/go/controllers/BudgetController";
import { useProjectStore } from "./stores/useProjectStore";
import { useGroupStore } from "./stores/useGroupStore";
import { useBudgetStore } from "./stores/useBudgetStore";
import { useAppStore } from "./stores/useAppStore";

const App: React.FC = () => {
  const [activeView, setActiveView] = useState<"home" | "projects" | "budgets" |  "groups" | "categories" | "analysis" |"transactions">("home");

  const appStore = useAppStore();

  return (
    <Theme>
    <div style={{ display: "flex", height: "100vh", width: "100%" }}>
      {/* Sidebar */}
      <nav
        style={{
          width: "200px",
          background: "#f4f4f4",
          padding: "20px",
          display: "flex",
          flexDirection: "column",   // <- vertical layout
          gap: "10px",               // spacing between buttons
          boxShadow: "2px 0 5px rgba(0,0,0,0.1)"
        }}
      >
        <button
          style={{
            padding: "10px",
            textAlign: "left",
            background: activeView === "home" ? "#ddd" : "#fff",
            border: "1px solid #ccc",
            borderRadius: "4px",
            cursor: "pointer"
          }}
          onClick={() => setActiveView("home")}
        >
          Home
        </button>
        <button
          style={{
            padding: "10px",
            textAlign: "left",
            background: activeView === "analysis" ? "#ddd" : "#fff",
            border: "1px solid #ccc",
            borderRadius: "4px",
            cursor: "pointer"
          }}
          onClick={() => setActiveView("analysis")}
        >
          Analysis
        </button>
        <button
          style={{
            padding: "10px",
            textAlign: "left",
            background: activeView === "projects" ? "#ddd" : "#fff",
            border: "1px solid #ccc",
            borderRadius: "4px",
            cursor: "pointer"
          }}
          onClick={() => setActiveView("projects")}
        >
          Projects
        </button>
        <button
          style={{
            padding: "10px",
            textAlign: "left",
            background: activeView === "budgets" ? "#ddd" : "#fff",
            border: "1px solid #ccc",
            borderRadius: "4px",
            cursor: "pointer"
          }}
          onClick={() => setActiveView("budgets")}
        >
          Budgets
        </button>
        <button
          style={{
            padding: "10px",
            textAlign: "left",
            background: activeView === "groups" ? "#ddd" : "#fff",
            border: "1px solid #ccc",
            borderRadius: "4px",
            cursor: "pointer"
          }}
          onClick={() => setActiveView("groups")}
        >
          Groups
        </button>
        <button
          style={{
            padding: "10px",
            textAlign: "left",
            background: activeView === "categories" ? "#ddd" : "#fff",
            border: "1px solid #ccc",
            borderRadius: "4px",
            cursor: "pointer"
          }}
          onClick={() => setActiveView("categories")}
        >
          Categories
        </button>
        <button
          style={{
            padding: "10px",
            textAlign: "left",
            background: activeView === "transactions" ? "#ddd" : "#fff",
            border: "1px solid #ccc",
            borderRadius: "4px",
            cursor: "pointer"
          }}
          onClick={() => setActiveView("transactions")}
        >
          Transactions
        </button>
      </nav>
      {/* Main content */}
      <main style={{ flex: 1, padding: "10px" }}>
          {/*{activeView === "transactions" && <Transactions projectStore={projectStore} budgetStore={null}/>}
          {activeView === "budgets" && <Budgets globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projects}/>}
        {activeView === "groups" && <Groups globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projects}/>}
        {activeView === "categories" && <Categories globalGroup={globalGroup} setGlobalGroup={setGlobalGroup} groupList={groups} globalBudget={globalBudget} setGlobalBudget={setGlobalBudget} budgetList={budgets} />}
        {activeView === "projects" && <Projects globalProject={globalProject} setGlobalProject={setGlobalProject} />}
        {activeView === "analysis" && <Analysis />}*/}
        {activeView === "home" && <Home appStore={appStore} />}
      </main>
    </div>
    </Theme>
  );
};

export default App;

