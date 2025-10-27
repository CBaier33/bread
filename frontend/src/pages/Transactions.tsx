import React, { useEffect, useState } from "react";
import { ListTransactions, CreateTransaction } from "../../wailsjs/go/controllers/TransactionController"
import { models } from "../../wailsjs/go/models";
import ProjectSelect from "../components/ProjectSelect";
import TransactionCard from "../components/TransactionCard";
import NewTransactionPrompt from "../components/NewTransactionPrompt";

interface TransactionProps {
  globalProject: models.Project;
  setGlobalProject: (project: models.Project) => void;
  projectList: models.Project[];
}

const Transactions: React.FC<TransactionProps> = ({ globalProject, setGlobalProject, projectList }) => {
  const [transactions, setTransactions] = useState<models.Transaction[]>([]);
  const [loading, setLoading] = useState(true);
  const [description, setDescription] = useState("");
  const [amount, setAmount] = useState<number>(0);
  const [date, setDate] = useState("");
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

  const handleCreateTransaction = async (description: string, amount: number, date: string, expenseType: boolean, notes: string, tags:string) => {
    try {
      await CreateTransaction(globalProject.id, null, description, amount, date, expenseType, notes);
      setDescription("");
      setDate("2003-05-02");
      setNotes("");
      setExpenseType(true);
      setAmount(0);
      loadTransactions();
    } catch (err) {
      console.error("Error adding transaction:", err);
    }

    console.log(tags);
  };

  if (loading) return <div>Loading transactions...</div>;


  return (
    <div className="p-5">
      <div className="mb-4 flex gap-2 items-center w-full">
        <h1 className="text-4xl font-bold">Transactions</h1>
      <ProjectSelect globalProject={globalProject} setGlobalProject={setGlobalProject} projectList={projectList}/>
        <NewTransactionPrompt onSave={handleCreateTransaction}/>
      </div>

      <div className="w-full gap-7">
        {transactions.map((p) => (
          <TransactionCard
            key={p.id}
            transaction={p}
            width=""
          />
        ))}
      </div>
    </div>
  );
};

export default Transactions;

