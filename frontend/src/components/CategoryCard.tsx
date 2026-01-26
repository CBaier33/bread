import { useEffect, useState } from "react";
import { Text } from "@radix-ui/themes";
import { models } from "../../wailsjs/go/models";
import { AllocationCost } from "../../wailsjs/go/controllers/AnalysisController";
import { useAppStore } from "../stores/useAppStore";

interface CategoryCardProps {
  category: models.Category;
  allocation?: models.BudgetAllocation;
  appStore: ReturnType<typeof useAppStore>;
}

const CategoryCard: React.FC<CategoryCardProps> = ({
  category,
  allocation,
  appStore,
}) => {
  const { selectedBudget } = appStore;

  const [actual, setActual] = useState<number>(0);
  const [loading, setLoading] = useState(false);

  
  useEffect(() => {
    const fetchActual = async () => {
      if (!selectedBudget?.id || !category?.id) {
        setActual(0);
        return;
      }

      setLoading(true);
      try {
        const result = await AllocationCost(
          selectedBudget.id,
          category.id
        );

        setActual(result ?? 0);
      } catch (err) {
        console.error("Failed to fetch allocation cost:", err);
        setActual(0);
      } finally {
        setLoading(false);
      }
    };

    fetchActual();
  }, [selectedBudget?.id, category?.id]);

  const expected = allocation?.expected_cost ?? 0;

  const formattedExpected = `$${(expected / 100).toFixed(2)}`;
  const formattedActual = `$${(actual / 100).toFixed(2)}`;

  const overBudget = actual > expected;

  return (
    <div className="flex flex-col items-center px-3 py-2 rounded-xl border bg-gray-50 text-gray-800 shadow-sm min-w-[7rem]">
      {/* Category Name */}
      <Text size="2" weight="medium" className="truncate">
        {category.name || category.description}
      </Text>

      {/* Expected */}
      <Text size="1" className="text-gray-500">
        Expected: {formattedExpected}
      </Text>

      {/* Actual */}
      {loading ? (
        <Text size="1" className="text-gray-400 animate-pulse">
          Actual: --
        </Text>
      ) : (
        <Text
          size="1"
          className={overBudget ? "text-red-600" : "text-green-700"}
        >
          Actual: {formattedActual}
        </Text>
      )}
    </div>
  );
};

export default CategoryCard;

