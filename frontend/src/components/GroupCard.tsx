import { useEffect, useState } from "react";
import { Text } from "@radix-ui/themes";
import { models } from "../../wailsjs/go/models";
import { ListCategories } from "../../wailsjs/go/controllers/CategoryController";
import { ListAllocations } from "../../wailsjs/go/controllers/BudgetController";
import { useAppStore } from "../stores/useAppStore";
import CategoryCard from "./CategoryCard";
import NewCategoryPrompt from "./NewCategoryPrompt";

interface GroupCardProps {
  group: models.Group;
  appStore: ReturnType<typeof useAppStore>;
}

const GroupCard: React.FC<GroupCardProps> = ({ group, appStore }) => {
  const [categories, setCategories] = useState<models.Category[]>([]);
  const [allocations, setAllocations] = useState<models.BudgetAllocation[]>([]);
  const [loadingCats, setLoadingCats] = useState(false);
  const [loadingAllocs, setLoadingAllocs] = useState(false);

  const { selectedBudget } = appStore;

  
  useEffect(() => {
    const fetchCategories = async () => {
      if (!group || group.id === 0) {
        setCategories([]);
        return;
      }

      setLoadingCats(true);
      try {
        const result = await ListCategories(group.id);
        setCategories(result ?? []);
      } catch (err) {
        console.error("Failed to fetch categories:", err);
        setCategories([]);
      } finally {
        setLoadingCats(false);
      }
    };

    fetchCategories();
  }, [appStore.selectedProject]);

  useEffect(() => {
    const fetchAllocations = async () => {
      if (!selectedBudget || selectedBudget.id === 0) {
        setAllocations([]);
        return;
      }

      setLoadingAllocs(true);
      try {
        const result = await ListAllocations(selectedBudget.id);
        setAllocations(result ?? []);
      } catch (err) {
        console.error("Failed to fetch allocations:", err);
        setAllocations([]);
      } finally {
        setLoadingAllocs(false);
      }
    };

    fetchAllocations();
  }, [selectedBudget?.id]);

  const groupName =
    group?.name && group.name.length > 0 ? group.name : "No Group Selected";

  return (
    <div className="w-full flex flex-col gap-3 items-center justify-center p-6 rounded-2xl border bg-white shadow-sm">

      {/* Group Title */}

      <div className="flex items-center gap-2">
      <Text
        size="7"
        weight="bold"
        className="text-violet-700 truncate max-w-full text-center"
      >
        {groupName}
      </Text>

      <NewCategoryPrompt group={group} appStore={appStore}/>
      </div>

      {/* Categories + Allocations */}
      <div className="w-full flex flex-wrap justify-center gap-2 mt-2">
        {loadingCats || loadingAllocs ? (
          <Text size="2" className="text-gray-400 animate-pulse">
            Loading...
          </Text>
        ) : categories.length === 0 ? (
          <Text size="2" className="text-gray-400 italic">
            No categories in this group
          </Text>
        ) : (
          categories.map((cat) => {
            const allocation = allocations.find(
              (a) => a.category_id === cat.id
            );

            return (
              <CategoryCard
                key={cat.id}
                category={cat}
                allocation={allocation}
                appStore={appStore}
              />
            );
          })
        )}
      </div>
    </div>
  );
};

export default GroupCard;

