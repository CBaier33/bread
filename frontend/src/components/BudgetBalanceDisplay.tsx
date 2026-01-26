import { useEffect, useState } from "react";
import { Text } from "@radix-ui/themes";
import { models } from "../../wailsjs/go/models";
import { BudgetBalance } from "../../wailsjs/go/controllers/AnalysisController";
import { useAppStore } from "../stores/useAppStore";

interface BudgetBalanceDisplayProps {
  appStore: ReturnType<typeof useAppStore>;
}

const BudgetBalanceDisplay: React.FC<BudgetBalanceDisplayProps> = ({ appStore, }) => {

  const [balance, setBalance] = useState<number>(0);
  const [loading, setLoading] = useState(false);

  const {
    selectedBudget: budget,
  } = appStore;

  useEffect(() => {
    const fetchBalance = async () => {
      if (!budget || budget.id === 0) {
        setBalance(0);
        return;
      }

      setLoading(true);

      try {
        const result = await BudgetBalance(budget?.id);

        setBalance(result ?? 0);
      } catch (err) {
        console.error("Failed to fetch budget balance:", err);
        setBalance(0);
      } finally {
        setLoading(false);
      }
    };

    fetchBalance();
  }, [budget]);

const formatted =
  `${balance < 0 ? "- " : ""}$${Math.abs(balance / 100).toFixed(2)}`;

  return (
    <div className="w-full flex flex-col items-center justify-center p-6 rounded-2xl border bg-white shadow-sm">
        <Text
          size="5"
          weight="bold"
        >
        Balance
        </Text>
      {loading ? (
        <Text size="5" className="animate-pulse text-gray-400">
          Loading...
        </Text>
      ) : (
        <Text
          size="8"
          weight="bold"
          className={`${
            balance < 0 ? "text-red-600" : "text-green-600"
          }`}
        >
          {formatted}
        </Text>
      )}
    </div>
  );
};

export default BudgetBalanceDisplay;

