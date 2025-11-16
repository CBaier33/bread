import React, { useEffect, useState } from "react";
import { ListTransactions, CreateTransaction, UpdateTransaction, DeleteTransaction } from "../../wailsjs/go/controllers/TransactionController"
import { models } from "../../wailsjs/go/models";
import ProjectSelect from "../components/ProjectSelect";
import TransactionCard from "../components/TransactionCard";
import NewTransactionPrompt from "../components/NewTransactionPrompt";
import { ListProjectCategories } from "../../wailsjs/go/controllers/CategoryController";

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
  const [date, setDate] = React.useState<Date | undefined>(new Date());
  const [notes, setNotes] = useState("");
  const [expenseType, setExpenseType] = useState(true);

  const [globalCategory, setGlobalCategory] = useState<models.Category>(new models.Category({
    id: 0, 
    description: "Empty",
  }));


  const [categories, setCategories] = useState<models.Category[]>([]);

  const fetchCategories = async () => {
    try {
      const result = await ListProjectCategories(globalProject.id);
      setCategories(result ?? []);
    } catch (err) {
      console.error("Failed to fetch projects:", err);
    }
  };

  const fetchTransactions = async () => {
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
    fetchTransactions();
    fetchCategories();
  }, [globalProject]);

  const handleCreateTransaction = async (description: string, amount: number, date: Date | undefined, expense: boolean, notes: string, tags:string) => {
    try {
      await CreateTransaction(globalProject.id, globalCategory.id, description, amount, (date ?? new Date()).toLocaleDateString(), expense, notes);
      setDescription("");
      setNotes("");
      setExpenseType(true);
      setAmount(0);
      fetchTransactions();
    } catch (err) {
      console.error("Error adding transaction:", err);
    }

    console.log(tags);
  };

  const handleEditTransaction = async (transaction: models.Transaction) => {
    try {
      await UpdateTransaction(transaction);
      await fetchTransactions();
    } catch (err) {
      console.error("Failed to create transaction:", err);
    }
  };

  const handleDeleteTransaction = async (transaction: models.Transaction) => {
    try {
      await DeleteTransaction(transaction.id);
      await fetchTransactions();
    } catch (err) {
      console.error("Failed to create transaction:", err);
    }
  };

  if (loading) return <div>Loading transactions...</div>;


  return (
    <div className="p-5 h-full flex flex-col">
      <div className="mb-4 flex gap-2 justify-between items-center">
        <h1 className="text-2xl font-bold">Transactions</h1>
        <NewTransactionPrompt onSave={handleCreateTransaction} globalCategory={globalCategory} setGlobalCategory={setGlobalCategory} categoryList={categories} />
      </div>

      {/* This is the scrollable list */}
      <div className="flex-1 overflow-auto p-2">
        {transactions.map((t) => (
          <TransactionCard 
            key={t.id} 
            transaction={t} 
            width="" 
            onEdit={handleEditTransaction}
            onDelete={handleDeleteTransaction}
            categoryList={categories}
          />
        ))}
      </div>
    </div>
  );
};

export default Transactions;

