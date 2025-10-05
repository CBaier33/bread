import React, { useState } from "react";
import { Dialog, Text, TextField, Flex, Button, IconButton } from "@radix-ui/themes";
import { PlusIcon } from "@radix-ui/react-icons";
import { models } from "../../wailsjs/go/models";

interface NewProjectPromptProps {
  onSave: (name: string, description: string, currency: string) => Promise<void> | void;
}

const NewProjectPrompt: React.FC<NewProjectPromptProps> = ({ onSave }) => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [currency, setCurrency] = useState("USD");

  const handleSave = async () => {
    await onSave(name, description, currency);

    setName("");
    setDescription("");
    setCurrency("USD");
  };

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <IconButton size="1" color="gray" radius="full" variant="solid">
          <PlusIcon />
        </IconButton>
      </Dialog.Trigger>

      <Flex direction="column" width="23rem" asChild>
        <Dialog.Content>
          <Dialog.Title>New Project</Dialog.Title>
          <Flex gap="2" direction="column">
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Name
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  value={name}
                  placeholder="Personal"
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
                  value={description}
                  placeholder="Capital One"
                  onChange={(e) => setDescription(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>

            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Currency
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  value={currency}
                  placeholder="USD"
                  onChange={(e) => setCurrency(e.target.value)}
                  disabled
                  size="2"
                />
              </Flex>
            </label>
              <Dialog.Close>
            <Flex gap="2" mt="2"> {/*///justify="end">*/}
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

export default NewProjectPrompt;
