import React, { useState, useEffect } from "react";
import { Dialog, Text, TextField, Flex, Button, SegmentedControl, Popover } from "@radix-ui/themes";
import { models } from "../../wailsjs/go/models";
import { DayPicker } from "react-day-picker";
import "react-day-picker/dist/style.css";
import CategorySelect from "./CategorySelect";
import { GetCategoryByID } from "../../wailsjs/go/controllers/CategoryController";

interface EditTransactionProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  transaction: models.Transaction;
  onSave: (updatedTransaction: models.Transaction) => Promise<void> | void;
  categoryList: models.Category[];
}

const EditTransaction: React.FC<EditTransactionProps> = ({ open, onOpenChange, transaction, onSave, categoryList }) => {
  const [description, setDescription] = useState(transaction.description);
  const [amount, setAmount] = useState(transaction.amount);
  const [displayAmount, setDisplayAmount] = useState((amount / 100).toFixed(2));
  const [expenseType, setExpenseType] = useState(transaction.expense_type);
  const [category, setCategory] = useState<models.Category>(new models.Category())
  const [date, setDate] = React.useState<Date | undefined>(new Date(transaction.date));
  const [notes, setNotes] = useState(transaction.notes);
  const [tags, setTags] = useState("");

  useEffect(() => {
    setDescription(transaction.description);
    setAmount(transaction.amount);
    const loadCategory = async () => {
      if (transaction.category_id) {
        const cat = await GetCategoryByID(transaction.category_id);
        setCategory(cat);
      } else {
        setCategory(new models.Category());
      }
    };

    loadCategory();
  }, [transaction]);

  const handleSave = async () => {
    await onSave({ 
      ...transaction, 
      description, 
      amount, 
      expense_type: expenseType, 
      date: (date ?? new Date()).toLocaleDateString(),
      category_id: category.id,
      notes: notes, 
    });

    setAmount(amount)
    onOpenChange(false);
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
    <Dialog.Root open={open} onOpenChange={onOpenChange}>
      <Dialog.Content maxWidth="480px">
        <Dialog.Title>Edit Transaction</Dialog.Title>

        <Flex direction="column" gap="3" mt="3">
          <label>
            <Text as="div" size="2" mb="1" weight="bold">
              Description
            </Text>
            <TextField.Root value={description} onChange={(e) => setDescription(e.target.value)} size="2" />
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
                globalCategory={category}
                setGlobalCategory={setCategory}
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

          <Flex justify="end" gap="3" mt="4">
            <Button variant="soft" color="gray" onClick={() => onOpenChange(false)}>
              Cancel
            </Button>
            <Button onClick={handleSave}>Save</Button>
          </Flex>
        </Flex>
      </Dialog.Content>
    </Dialog.Root>
  );
};

export default EditTransaction;
