import React, { useEffect, useState } from "react";
import { ListTransactions, AddTransaction } from "../../wailsjs/go/controllers/TransactionController"
import { models } from "../../wailsjs/go/models";

const Transactions: React.FC = () => {
  const [transactions, setTransactions] = useState<models.Transaction[]>([]);
  const [loading, setLoading] = useState(true);
  const [budgetID, setBudgetID] = useState<number>(1);
  const [description, setDescription] = useState("");
  const [amount, setAmount] = useState<number>(0);
  const [date, setDate] = useState("2003-05-02");
  const [notes, setNotes] = useState("");

  const loadTransactions = async () => {
    try {
      const txs = await ListTransactions(null);
      setTransactions(txs ?? []); // types now match
    } catch (err) {
      console.error("Error fetching transactions:", err);
  } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadTransactions();
  }, []);

  const handleAdd = async () => {
    try {
      await AddTransaction(description, amount, date, notes, budgetID, null, null);
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
    <div>
      <h1>Transactions</h1>
      <ul>
        {transactions.map((tx) => (
          <li key={tx.id}>
            {tx.description} - ${tx.amount} - ${tx.date} - ${tx.category_name} - {tx.created_at}
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

