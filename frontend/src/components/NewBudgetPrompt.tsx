import React, { useState } from "react";
import { Dialog, Popover, Text, TextField, Flex, Button, IconButton, SegmentedControl } from "@radix-ui/themes";
import { PlusIcon } from "@radix-ui/react-icons";
import { DayPicker } from "react-day-picker";
import "react-day-picker/dist/style.css";
import dayjs from "dayjs";
import { models } from "../../wailsjs/go/models";

interface NewBudgetPromptProps {
  onSave: (name: string, startDate: Date | undefined, endDate: Date | undefined, expectedIncome: number, startingBalance: number) => Promise<void> | void;
}

const NewBudgetPrompt: React.FC<NewBudgetPromptProps> = ({ onSave }) => {
  const [name, setName] = useState("");
  const [startDate, setStartDate] = React.useState<Date>();
  const [endDate, setEndDate] = React.useState<Date>();
  const [expectedIncome, setExpectedIncome] = useState<number>(0);
  const [displayExpected, setDisplayExpected] = useState((0 / 100).toFixed(2));
  const [startingBalance, setStartingBalance] = useState<number>(0);
  const [displayStarting, setDisplayStarting] = useState((0 / 100).toFixed(2));


  const handleCreateBudget = async () => {
    console.log("handeCreateBudget", startDate, endDate)
    await onSave(name, startDate, endDate, expectedIncome, startingBalance);
  };

  const handleExpected = async (value: string) => {
    setDisplayExpected(value)
    const parsed = parseFloat(value)
    if (!isNaN(parsed)) {
      setExpectedIncome(Math.round(parsed * 100));
    }
  }

  const handleStarting = async (value: string) => {
    setDisplayStarting(value)
    const parsed = parseFloat(value)
    if (!isNaN(parsed)) {
      setStartingBalance(Math.round(parsed * 100));
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
          <Dialog.Title>New Budget</Dialog.Title>
          <Flex gap="2" direction="column">
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Name
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  defaultValue="Monthly Budget"
                  value={name}
                  placeholder="Monthly Budget"
                  onChange={(e) => setName(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
              <label>
                <Text as="div" size="2" mt="1" weight="bold">
                  Date Range
                </Text>
              </label>
            <Popover.Root>
              <Flex width="20rem" asChild>
                    <Popover.Trigger>
                      <TextField.Root defaultValue={startDate ? startDate.toLocaleDateString() : "Select start"}>
                      </TextField.Root>
                    </Popover.Trigger>
              </Flex>
              <Popover.Content size="1" align="start" >
                  <DayPicker
                    mode="single"
                    selected={startDate}
                    onSelect={setStartDate}
                    numberOfMonths={1}
                  />

                  <Popover.Close>
                    <Flex justify="end">
                      <Button>Select</Button>
                    </Flex>
                  </Popover.Close>
              </Popover.Content>
            </Popover.Root>
            <Popover.Root>
              <Flex width="20rem" asChild>
                    <Popover.Trigger>
                      <TextField.Root defaultValue={endDate ? endDate.toLocaleDateString() : "Select end"}>
                      </TextField.Root>
                    </Popover.Trigger>
              </Flex>
              <Popover.Content size="1" align="end" >
                  <DayPicker
                    mode="single"
                    selected={endDate}
                    onSelect={setEndDate}
                    numberOfMonths={1}
                  />

                  <Popover.Close>
                    <Flex justify="end">
                      <Button>Select</Button>
                    </Flex>
                  </Popover.Close>
              </Popover.Content>
            </Popover.Root>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Expected Income
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  type="number"
                  step="0.01"
                  value={displayExpected}
                  onChange={(e) => handleExpected(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Starting Balance
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  type="number"
                  step="0.01"
                  value={displayStarting}
                  onChange={(e) => handleStarting(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
          </Flex>
            <Dialog.Close>
              <Flex gap="3" mt="4" justify="end">
                  <Button variant="soft" color="gray">
                    Cancel
                  </Button>
                  <Button onClick={handleCreateBudget}>
                    Save
                  </Button>
              </Flex>
            </Dialog.Close>
        </Dialog.Content>
        </Flex>

      </Dialog.Root>
  );
};

export default NewBudgetPrompt;
