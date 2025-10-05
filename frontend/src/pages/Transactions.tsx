import React, { useEffect, useState } from "react";
import { ListTransactions, CreateTransaction } from "../../wailsjs/go/controllers/TransactionController"
import { models } from "../../wailsjs/go/models";
import ProjectSelect from "../components/ProjectSelect";

interface TransactionProps {
  globalProject: models.Project;
  setGlobalProject: (project: models.Project) => void;
  projectList: models.Project[];
}

const Transactions: React.FC<TransactionProps> = ({ globalProject, setGlobalProject, projectList }) => {
  const [transactions, setTransactions] = useState<models.Transaction[]>([]);
  const [loading, setLoading] = useState(true);
  const [budgetID, setBudgetID] = useState<number>(0);
  const [description, setDescription] = useState("");
  const [amount, setAmount] = useState<number>(0);
  const [date, setDate] = useState("2003-05-02");
  const [notes, setNotes] = useState("");
  const [expenseType, setExpenseType] = useState(true);

  const loadTransactions = async () => {
    try {
      const txs = await ListTransactions(globalProject.id, null, null);
      setTransactions(txs ?? []); // types now match
    } catch (err) {
      console.error("Error fetching transactions:", err);
  } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadTransactions();
  }, [globalProject]);

  const handleAdd = async () => {
    try {
      await CreateTransaction(globalProject.id, null, description, amount, date, expenseType, notes, );
      setDescription("");
      setDate("2003-05-02");
      setNotes("")
      setAmount(0);
      loadTransactions();
    } catch (err) {
      console.error("Error adding transaction:", err);
    }
  };

  if (loading) return <div>Loading transactions...</div>;

  return (
    <div className="p-4">
      <h1>Transactions</h1>
      <h2>Current Project: {globalProject.name}</h2>
      <ProjectSelect globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projectList}/>
      <ul>
        {transactions.map((tx) => (
          <li key={tx.id}>
            {tx.description} - ${tx.amount} - ${tx.date} - {tx.created_at}
          </li>
        ))}
      </ul>
      <div>
        <input
          type="text"
          placeholder="Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <input
          type="number"
          placeholder="Amount"
          value={amount}
          onChange={(e) => setAmount(parseFloat(e.target.value))}
        />
        <input
          type="text"
          placeholder="Date"
          value={date}
          onChange={(e) => setDate(e.target.value)}
        />
        <input
          type="text"
          placeholder="Notes"
          value={notes}
          onChange={(e) => setNotes(e.target.value)}
        />
        <button onClick={handleAdd}>Add Transaction</button>
      </div>
    </div>
  );
};

export default Transactions;

