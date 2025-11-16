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

const App: React.FC = () => {
  const [activeView, setActiveView] = useState<"home" | "projects" | "budgets" |  "groups" | "categories" | "analysis" |"transactions">("home");

  const [projects, setProjects] = useState<models.Project[]>([]); 
  const [globalProject, setGlobalProject] = useState<models.Project>(new models.Project({
    id: 0, 
    name: "Empty",
    description: "Empty",
    currency: "Empty",
  }));

  const [groups, setGroups] = useState<models.Group[]>([]); 
  const [globalGroup, setGlobalGroup] = useState<models.Group>(new models.Group({
    id: 0, 
    description: "Empty",
  }));

  const [budgets, setBudgets] = useState<models.Budget[]>([]); 
  const [globalBudget, setGlobalBudget] = useState<models.Budget>(new models.Budget({
    id: 0, 
    description: "Empty",
  }));

  useEffect(() => {
  fetchProjects();
  if (projects.length > 0 && globalProject.id == 0) {
    setGlobalProject(projects[0]);
  };
  fetchGroups();
  if (groups.length > 0 && globalGroup.id == 0) {
    setGlobalGroup(groups[0]);
  }
}, [projects, groups]);

  const fetchProjects = async () => {
    try {
      const result = await ListProjects();
      setProjects(result ?? []);
    } catch (err) {
      console.error("Failed to fetch projects:", err);
    }
  };

  const fetchGroups = async () => {
    try {
      const result = await ListGroups(globalProject.id);
      setGroups(result ?? []);
    } catch (err) {
      console.error("Failed to fetch projects:", err);
    }
  };

  const fetchBudgets = async () => {
    try {
      const result = await ListBudgets(globalProject.id);
      setBudgets(result ?? []);
    } catch (err) {
      console.error("Failed to fetch projects:", err);
    }
  };

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
        {activeView === "transactions" && <Transactions globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projects}/>}
        {activeView === "budgets" && <Budgets globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projects}/>}
        {activeView === "groups" && <Groups globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projects}/>}
        {activeView === "categories" && <Categories globalGroup={globalGroup} setGlobalGroup={setGlobalGroup} groupList={groups} globalBudget={globalBudget} setGlobalBudget={setGlobalBudget} budgetList={budgets} />}
        {activeView === "projects" && <Projects globalProject={globalProject} setGlobalProject={setGlobalProject} />}
        {activeView === "analysis" && <Analysis />}
        {activeView === "home" && <Home 
            globalProject={globalProject} setGlobalProject={setGlobalProject} projects={projects}
            globalGroup={globalGroup} setGlobalGroup={setGlobalGroup} groupList={groups}
            globalBudget={globalBudget} setGlobalBudget={setGlobalBudget} budgetList={budgets}
          />}
      </main>
    </div>
    </Theme>
  );
};

export default App;

