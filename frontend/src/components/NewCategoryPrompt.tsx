import React, { useState } from "react";
import { Dialog, Popover, Text, TextField, Flex, Button, IconButton, SegmentedControl } from "@radix-ui/themes";
import { PlusIcon } from "@radix-ui/react-icons";
import "react-day-picker/dist/style.css";
import { models } from "../../wailsjs/go/models";
import BudgetSelect from "./BudgetSelect";

interface NewCategoryPromptProps {
  onSave: (name: string, description: string, expenseType: boolean, expected: number) => Promise<void> | void;
  globalBudget: models.Budget;
  setGlobalBudget: (group: models.Budget) => void;
  budgetList: models.Budget[];
}

const NewCategoryPrompt: React.FC<NewCategoryPromptProps> = ({ onSave, globalBudget, setGlobalBudget, budgetList }) => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [expenseType, setExpenseType] = useState(true);
  const [amount, setAmount] = useState(0);
  const [displayAmount, setDisplayAmount] = useState((0 / 100).toFixed(2));

  const handleCreateCategory = async () => {
    await onSave(name, description, expenseType, amount);
    setDescription("");
    setAmount(0);
    setDisplayAmount((0 / 100).toFixed(2))
  };

  const handleAmount = async (value: string) => {
    setDisplayAmount(value)
    const parsed = parseFloat(value)
    if (!isNaN(parsed)) {
      setAmount(Math.round(parsed * 100));
    }
  }

  const handleExpense= async (value: string) => {
    if (value == "deposit") {
      setExpenseType(false)
    } else if (value == "withdrawl") {
      setExpenseType(true)
    }
  }

  return (
      <Dialog.Root>
        <Dialog.Trigger>
          <IconButton size="1" color="gray" radius="full" variant="solid">
            <PlusIcon />
          </IconButton>
        </Dialog.Trigger>
        <Flex direction="column" width="32rem" asChild>
        <Dialog.Content>
          <Dialog.Title>New Category</Dialog.Title>
          <Flex gap="2" direction="column">
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Name
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  defaultValue=""
                  value={name}
                  placeholder=""
                  onChange={(e) => setName(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Description
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  defaultValue=""
                  value={description}
                  placeholder=""
                  onChange={(e) => setDescription(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Amount
              </Text>
              <Flex width="20rem" gap="2">
                <TextField.Root
                  type="number"
                  step="0.01"
                  value={displayAmount}
                  onChange={(e) => handleAmount(e.target.value)}
                  size="2"
                />
                <SegmentedControl.Root 
                  defaultValue="withdrawl"
                  onValueChange={handleExpense}>
                  <SegmentedControl.Item value="withdrawl">Withdrawl</SegmentedControl.Item>
                  <SegmentedControl.Item value="depost">Deposit</SegmentedControl.Item>
                </SegmentedControl.Root>
              </Flex>
            </label>
            <BudgetSelect globalBudget={globalBudget} setGlobalBudget={setGlobalBudget} budgetList={budgetList}/>
          </Flex>
            <Dialog.Close>
              <Flex gap="3" mt="4" justify="end">
                  <Button variant="soft" color="gray">
                    Cancel
                  </Button>
                  <Button onClick={handleCreateCategory}>
                    Save
                  </Button>
              </Flex>
            </Dialog.Close>
        </Dialog.Content>
        </Flex>

      </Dialog.Root>
  );
};

export default NewCategoryPrompt;
