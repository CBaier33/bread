import React, { useState } from "react";
import { Theme } from "@radix-ui/themes"
import Transactions from "./components/Transactions";
import Budgets from "./components/Budgets"
//import Home from "./Home";

const App: React.FC = () => {
  const [activeView, setActiveView] = useState<"home" | "budgets" | "transactions">("transactions");

  return (
    <Theme>
    <div style={{ display: "flex", height: "100vh" }}>
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
      <main style={{ flex: 1, padding: "20px" }}>
        {activeView === "transactions" && <Transactions />}
        {activeView === "budgets" && <Budgets />}
      </main>
    </div>
    </Theme>
  );
};

export default App;

