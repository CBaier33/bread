import React, { useState } from "react";
import { Dialog, Popover, Text, TextField, Flex, Button, IconButton, SegmentedControl } from "@radix-ui/themes";
import { PlusIcon } from "@radix-ui/react-icons";
import { DayPicker } from "react-day-picker";
import "react-day-picker/dist/style.css";
import { models } from "../../wailsjs/go/models";
import CategorySelect from "./CategorySelect";

interface NewTransactionPromptProps {
  onSave: (description: string, amount: number, date: Date | undefined, expenseType: boolean, notes: string, tags: string) => Promise<void> | void;
  globalCategory: models.Category;
  setGlobalCategory: (project: models.Category) => void;
  categoryList: models.Category[];
}

const NewTransactionPrompt: React.FC<NewTransactionPromptProps> = ({ onSave, globalCategory, setGlobalCategory, categoryList }) => {
  const [description, setDescription] = useState("");
  const [amount, setAmount] = useState(0);
  const [displayAmount, setDisplayAmount] = useState((0 / 100).toFixed(2));
  const [expenseType, setExpenseType] = useState(true);
  const [date, setDate] = React.useState<Date | undefined>();
  const [notes, setNotes] = useState("");
  const [tags, setTags] = useState("");

  const handleSave = async () => {
    await onSave(description, amount, date, expenseType, notes, tags);

    setDescription("");
    setDate(new Date())
    setAmount(0);
    setDisplayAmount((0 / 100).toFixed(2))
    setExpenseType(true)
    setNotes("")
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

      <Flex direction="column" width="23rem" asChild>
        <Dialog.Content>
          <Dialog.Title>New Transaction</Dialog.Title>
          <Flex gap="2" direction="column">
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Description
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
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
                  defaultValue={expenseType ? "withdrawl" : "deposit"}
                  onValueChange={handleExpense}>
                  <SegmentedControl.Item value="withdrawl">Withdrawl</SegmentedControl.Item>
                  <SegmentedControl.Item value="deposit">Deposit</SegmentedControl.Item>
                </SegmentedControl.Root>
              </Flex>
            </label>
            <label>
            <Text as="div" size="2" mb="1" weight="bold">
              Date
            </Text>
            <Popover.Root>
              <Flex width="20rem" asChild>
                    <Popover.Trigger>
                      <TextField.Root defaultValue={date ? date.toLocaleDateString() : "Select date"}>
                      </TextField.Root>
                    </Popover.Trigger>
              </Flex>
              <Popover.Content size="1" align="start" >
                  <DayPicker
                    mode="single"
                    selected={date}
                    onSelect={setDate}
                  />

                  <Popover.Close>
                    <Flex justify="end">
                      <Button>Select</Button>
                    </Flex>
                  </Popover.Close>
              </Popover.Content>
            </Popover.Root>
            </label>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Category
              </Text>
              <Flex width="20rem" asChild>
                <CategorySelect 
                  globalCategory={globalCategory}
                  setGlobalCategory={setGlobalCategory}
                  categoryList={categoryList}
                />
              </Flex>
            </label>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Notes
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  value={notes}
                  placeholder=""
                  onChange={(e) => setNotes(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Tags
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  value={tags}
                  placeholder="TBD"
                  onChange={(e) => setTags(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>

              <Dialog.Close>
            <Flex gap="2" mt="2"> 
                  <Button onClick={handleSave}>Save</Button>
                  <Button variant="soft" color="gray">
                    Cancel
                  </Button>
            </Flex>
              </Dialog.Close>
          </Flex>
        </Dialog.Content>
      </Flex>
    </Dialog.Root>
  );
};

export default NewTransactionPrompt;
