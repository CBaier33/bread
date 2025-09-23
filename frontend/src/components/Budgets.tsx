import React, { useEffect, useState } from "react";
import { Button, Popover, TextArea, IconButton, Dialog, Text, TextField, Box, Flex, Avatar, Checkbox } from "@radix-ui/themes";
import {PlusIcon, ChatBubbleIcon} from "@radix-ui/react-icons"
import { DayPicker } from "react-day-picker";
import "react-day-picker/dist/style.css";
import "./budget.css"
import dayjs from "dayjs";

import { ListBudgets, CreateBudget } from "../../wailsjs/go/controllers/BudgetController";
import { models } from "../../wailsjs/go/models";

const Budgets: React.FC = () => {
  const [budgets, setBudgets] = useState<models.Budget[]>([]);
  const [name, setName] = useState("");
  const [startDate, setStartDate] = React.useState<Date | undefined>();
  const [endDate, setEndDate] = React.useState<Date | undefined>();

  // Fetch budgets from backend
  const fetchBudgets = async () => {
    try {
      const result = await ListBudgets();
      console.log(result);
      setBudgets(result ?? []);
    } catch (err) {
      console.error("Failed to fetch budgets:", err);
    }
  };

  useEffect(() => {
    fetchBudgets();
  }, []);

  function formatDate(date?: Date): string {
    if (!date) return "";
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0"); // months are 0-indexed
    const day = String(date.getDate()).padStart(2, "0");
    return `${year}-${month}-${day}`;
}


  const handleCreateBudget = async () => {
    const start = formatDate(startDate) ?? ""
    const end = formatDate(endDate) ?? "";

    try {
      await CreateBudget(name, start, end);
      fetchBudgets();

      // Clear inputs
      setName("");
      setStartDate(undefined);
      setEndDate(undefined);

      // Refresh list
      fetchBudgets();
    } catch (err) {
      console.error("Failed to create budget:", err);
    }
  };

  return (
    <div className="p-4 max-w-md mx-auto">
      <h2 className="text-4xl font-bold mb-4">Budgets</h2>


      {/* Budget List */}
      <ul className="mb-6">
        {budgets.map((b) => (
          <li key={b.id} className="mb-2 p-2 border rounded">
            <strong>{b.name}</strong> (
            {dayjs(b.period_start).format("YYYY-MM-DD")} - {dayjs(b.period_end).format("YYYY-MM-DD")})
          </li>
        ))}
      </ul>

      
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

      {/*<Dialog.Root>
        <Dialog.Trigger>
          <IconButton size="1" color="gray" radius="full" variant="solid">
            <PlusIcon />
          </IconButton>
        </Dialog.Trigger>

        <Dialog.Content>
          <Dialog.Title mb="4">New Budget</Dialog.Title>
          <Dialog.Description></Dialog.Description>
          <Flex direction="column" gap="2" width="22rem">

                <label>
                  <Text as="div" size="2" mb="1" weight="bold">
                    Name
                  </Text>
                  <TextField.Root
                    defaultValue="Monthly Budget"
                    value={name}
                    placeholder="Enter a name for your budget"
                    onChange={(e) => setName(e.target.value)}
                  />
                </label>
                <label>
                  <Text as="div" size="2" mt="2" mb="2" weight="bold">
                    Date Range
                  </Text>
                </label>

             <Popover.Root>
                <Popover.Trigger>
                  <TextField.Root defaultValue={startDate ? startDate.toLocaleDateString() : "Select start date"}>
                  </TextField.Root>
                </Popover.Trigger>


                <Popover.Content
                  align="start"
                >
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
                <Popover.Trigger>
                  <TextField.Root defaultValue={endDate ? endDate.toLocaleDateString() : "Select end date"}>
                  </TextField.Root>
                </Popover.Trigger>

                <Popover.Content
                  align="start"
                  className="rounded-lg border bg-gray"
                  width="10px"
               >
                  <DayPicker
                    mode="single"
                    selected={endDate}
                    onSelect={setEndDate}
                    numberOfMonths={1}
                  />

                <Flex justify="end">
                  <Popover.Close>
                      <Button>Select</Button>
                  </Popover.Close>
                </Flex> 

                </Popover.Content>
              </Popover.Root>
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

      {/* Create Budget Dialog 
      </Dialog.Root>*/}
    </div>
  );
};

export default Budgets;

