import React, { useState, useEffect } from "react";
import { Theme } from "@radix-ui/themes"
import { models } from "../wailsjs/go/models";
import { ListProjects } from "../wailsjs/go/controllers/ProjectController";
import Transactions from "./pages/Transactions";
import Projects from "./pages/Projects"
import Budgets from "./pages/Budgets"
import Analysis from "./pages/Analysis";
import Home from "./pages/Home";

const App: React.FC = () => {
  const [activeView, setActiveView] = useState<"home" | "projects" | "budgets" | "analysis" |"transactions">("home");

  const [projects, setProjects] = useState<models.Project[]>([]); 
  const [globalProject, setGlobalProject] = useState<models.Project>(new models.Project({
    id: 0, 
    name: "Empty",
    description: "Empty",
    currency: "Empty",
  }));

  useEffect(() => {
  fetchProjects();
  if (projects.length > 0 && globalProject.id == 0) {
    setGlobalProject(projects[0]);
  }
}, [projects]);

  const fetchProjects = async () => {
    try {
      const result = await ListProjects();
      setProjects(result ?? []);
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
        {activeView === "budgets" && <Budgets />}
        {activeView === "projects" && <Projects globalProject={globalProject} setGlobalProject={setGlobalProject} />}
        {activeView === "analysis" && <Analysis />}
        {activeView === "home" && <Home />}
      </main>
    </div>
    </Theme>
  );
};

export default App;

