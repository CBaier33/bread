import { useState } from "react";

export interface GlobalStore<T> {
  items: T[];
  setItems: (items: T[]) => void;
  selected: T;
  setSelected: (item: T) => void;
  fetch: () => Promise<void>;
}

//export interface GlobalStore<T> {
//  items: T[];
//  setItems: React.Dispatch<React.SetStateAction<T[]>>;
//  selected: T | null;
//  setSelected: React.Dispatch<React.SetStateAction<T | null>>;
//  fetch: () => Promise<void>;
//}


//export function createGlobalStore<T>(
//  fetchFn: () => Promise<T[]>,
//  emptyItem: T
//) {
//  return function useStore(): GlobalStore<T> {
//    const [items, setItems] = useState<T[]>([]);
//    const [selected, setSelected] = useState<T | null>(null);
//
//    const fetch = async () => {
//      const result = await fetchFn();
//      const safeResult = result ?? [];
//
//      setItems(safeResult);
//
//      // ✅ Only auto-select if nothing is currently selected
//      if (safeResult.length > 0 && selected === null) {
//        setSelected(safeResult[0]);
//      }
//
//      // ✅ Clear selection if list is empty
//      if (safeResult.length === 0) {
//        setSelected(null);
//      }
//    };
//
//    return {
//      items,
//      setItems,
//      selected,
//      setSelected,
//      fetch,
//    };
//  };
//}

export function createGlobalStore<T>(
  fetchFn: () => Promise<T[]>,
  emptyItem: T
) {
  return function useStore(): GlobalStore<T> {
    const [items, setItems] = useState<T[]>([]);
    const [selected, setSelected] = useState<T>(emptyItem);

    const fetch = async () => {
      const result = await fetchFn();
      setItems(result ?? []);

      if (result && result.length > 0 && selected === emptyItem) {
        setSelected(result[0]);
      }
    };

    return {
      items,
      setItems,
      selected,
      setSelected,
      fetch,
    };
  };
}

